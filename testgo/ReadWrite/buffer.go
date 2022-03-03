package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	log.SetFlags(log.Lshortfile)

	/**
	bytes.Buffer
	strings.Builder
	*/

	/**
	将缓存内容写入文件
	*/
	create, err := os.Create("/data/home/xianlin_yang/Downloads/tmp.log")
	if err != nil {
		log.Fatalln(err)
	}
	buffer0 := bytes.NewBufferString("this is bytes.NewBufferString")
	//to, err := buffer0.WriteTo(create)
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(to)
	fprintf, err := fmt.Fprintf(create, buffer0.String())
	if err != nil {
		log.Fatalln(err)
	}

	by, _ := ioutil.ReadAll(create)
	fmt.Println(string(by))
	log.Println(fprintf)

	//遇到大字符串拼接时,也比较推荐这种形式
	buffer := bytes.Buffer{}
	buffer.WriteString("你好")
	buffer.WriteString("你是")
	buffer.WriteString("哪里")
	log.Println(buffer.String()) //你好你是哪里

	builder := strings.Builder{}
	builder.WriteString("你好")
	builder.WriteString("我是")
	builder.WriteString("地球")
	log.Println(builder.String()) //你好我是地球

	/**
	1. 合并字符串
	2. 读取前三个字节,存入到s3中
	3. 再次读取前三个字节,存入到s4中
	4. 打印s3和s4
	结果: 从字符串中读取3个字节的数据,再次打印时,发现前三个已消失,可见buffer中的数据通过read只能被消费一次
	*/
	//合并两个字符串
	s1 := "simple there"
	s2 := " yes copy"
	buffer1 := bytes.NewBufferString(s1)
	buffer1.WriteString(s2)
	log.Println(buffer1.String()) //simple there yes copy
	//从字符串中读取3个字节的数据,再次打印时,发现前三个已消失,可见buffer中的数据只能被消费一次
	s3 := make([]byte, 3)
	buffer1.Read(s3)
	log.Println(buffer1.String()) // lo there yes copy
	s4 := make([]byte, 3)
	buffer1.Read(s4)
	log.Println(buffer1.String()) // there yes copy
	//打印s3和s4
	log.Println(string(s3)) //hel
	log.Println(string(s4)) //lo

	/**
	读取一个字节,同样的一个直接只能被消费一次
	*/
	buffer2 := bytes.NewBufferString("this is buffer2")
	log.Println(buffer2.String()) // this is s5
	readByte, err := buffer2.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(readByte)) //t
	log.Println(buffer2.String()) //   his is s5

	/**
	读取字节到指定的分隔符,同样的一个直接只能被消费一次
	*/
	buffer3 := bytes.NewBufferString("this is buffer3")
	log.Println(buffer3.String()) // this is buffer3
	//读取到指定的符号位置,这里指定符号为空格
	readBytes, err := buffer3.ReadBytes(byte(' ')) //读取到第一个空格的位置
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(readBytes)) //this
	log.Println(buffer3.String())  //  is buffer3

	/**
	向后读取指定的字节数,同样的一个直接只能被消费一次
	*/
	buffer4 := bytes.NewBufferString("this is buffer4")
	log.Println(buffer4.String())
	a := buffer4.Next(2)
	log.Println(string(a))        //th
	log.Println(buffer4.String()) // is is buffer4
	b := buffer4.Next(2)
	log.Println(string(b))        //is
	log.Println(buffer4.String()) //is buffer4

	/**
	读取字符数
	*/
	buffer5 := bytes.NewBufferString("这里就是 buffer5")
	log.Println(buffer5.String()) // 这里 就是 buffer5
	//读取字符一直到指定的符号处
	readString, err := buffer5.ReadString(' ')
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(readString) //这里就是
	//读取一个UTF8的字符串
	buffer6 := bytes.NewBufferString("这里就是 simple buffer6")
	readRune, _, err := buffer6.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(readRune)         //36825
	log.Println(string(readRune)) //这

	/**
	撤销最后一次读取的数据
	*/
	//撤消最后一次读出的字节
	buffer7 := bytes.NewBufferString("this is buffer7")
	log.Println(buffer7.String()) //this is buffer7
	rs, err := buffer7.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(rs))       //t
	log.Println(buffer7.String()) // his is buffer7
	//撤消最后一次读出的字节
	err = buffer7.UnreadByte()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(buffer7.String()) // this is buffer7

	//撤消最后一次读出的 Unicode 字符
	buffer8 := bytes.NewBufferString("这里是 is buffer8")
	log.Println(buffer8.String()) // 这里是 is buffer8
	rb, _, err := buffer8.ReadRune()
	if err != nil {
		log.Fatalln(string(rb))
	}
	log.Println(buffer8.String()) // 里是 is buffer8
	//撤消最后一次读出的 Unicode 字符
	err = buffer8.UnreadRune()
	log.Println(buffer8.String()) //  这里是 is buffer8

}
