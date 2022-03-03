package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
)

const fileChunk = 8192 // 8KB

func countFileMd5(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return err.Error()
	}
	defer file.Close()

	info, _ := file.Stat()
	fileSize := info.Size()

	blocks := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		buf := make([]byte, blockSize)

		file.Read(buf)
		io.WriteString(hash, string(buf))
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}


func md5test2(){
	//rfile,_ := ioutil.ReadFile("/home/xianlin_yang/Desktop/nodetest/fileTest/kad/a.txt")
    r,err := os.Open("/home/xianlin_yang/Desktop/nodetest/fileTest/kad/a.txt")
	if err != nil{
		fmt.Println(err.Error())
	}

	info,_ := r.Stat()
	buf := make([]byte,info.Size())
	r.Read(buf)
	hash := md5.New()
	n,err := io.WriteString(hash,string(buf))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(n)
	fmt.Println(fmt.Sprintf("%x",hash.Sum(nil)))

}
func main() {
	fmt.Println(countFileMd5("/home/xianlin_yang/Desktop/nodetest/fileTest/kad/2.jpg"))
	md5test2()
}
