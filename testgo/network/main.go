

package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

/* 在windows下无法编译linux上的程序,主要是pcap在c++层次上的实现和windows不一样,需要在linux上安装go环境和libpcap
 * yum install libpcap
 * yum install libpcap-devel
 */
var (
	downStreamDataSize = 0 // 单位时间内下行的总字节数
	upStreamDataSize   = 0 // 单位时间内上行的总字节数
	//deviceName         = flag.String("i", "eth0", "network interface device name") // 要监控的网卡名称
)

func main() {
	//flag.Parse()

	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Printf("error:%v\n", err)
		log.Fatal(err)
	}
	fmt.Println("devices:",devices)
	// Find exact device
	var device pcap.Interface
	var ip, macAddr string = "", ""
	for _, d := range devices {
		ip, err = findDeviceIpv4(d)
		if err != nil {
			fmt.Printf("find device ip fail:%s\n", err.Error())
		} else {
			macAddr, err = findMacAddrByIp(ip)
			if err != nil {
				fmt.Printf("find mac fail:%s\n", err.Error())
			} else {
				fmt.Printf("ip:%s,mac:%s\n", ip, macAddr)
				if macAddr != "" {
					device = d
					break
				}
			}
		}
	}

	// 获取网卡handler，可用于读取或写入数据包
	handle, err := pcap.OpenLive(device.Name, 1024 /*每个数据包读取的最大值*/, true /*是否开启混杂模式*/, 30*time.Second /*读包超时时长*/)
	if err != nil {
		fmt.Printf("pcap open live error:%s\n", err.Error())
		return
	}
	defer handle.Close()

	// 开启子线程，每一秒计算一次该秒内的数据包大小平均值，并将下载、上传总量置零
	go monitor()

	// 开始抓包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// 只获取以太网帧
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			ethernet := ethernetLayer.(*layers.Ethernet)
			// 如果封包的目的MAC是本机则表示是下行的数据包,否则为上行
			if ethernet.DstMAC.String() == macAddr {
				downStreamDataSize += len(packet.Data()) // 统计下行封包总大小
			} else {
				upStreamDataSize += len(packet.Data()) // 统计上行封包总大小
			}
		}
	}
	return
}

// 获取网卡的IPv4地址
func findDeviceIpv4(device pcap.Interface) (ip string, err error) {
	for _, addr := range device.Addresses {
		if ipv4 := addr.IP.To4(); ipv4 != nil {
			ip = ipv4.String()
			return
		}
	}
	err = errors.New("device has no IPv4")
	return
}

// 根据网卡的IPv4地址获取MAC地址
// 有此方法是因为gopacket内部未封装获取MAC地址的方法,所以这里通过找到IPv4地址相同的网卡来寻找MAC地址
func findMacAddrByIp(ip string) (mac string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			break
		}

		for _, addr := range addrs {
			if a, ok := addr.(*net.IPNet); ok {
				if ip == a.IP.String() {
					mac = i.HardwareAddr.String()
					break
				}
			}
		}
	}
	return
}

// 每一秒计算一次该秒内的数据包大小平均值，并将下载、上传总量置零
func monitor() {
	for {
		os.Stdout.WriteString(fmt.Sprintf("\rDown:%.2fkb/s \t Up:%.2fkb/s", float32(downStreamDataSize)/1024/1, float32(upStreamDataSize)/1024/1))
		downStreamDataSize = 0
		upStreamDataSize = 0
		time.Sleep(1 * time.Second)
	}
}
