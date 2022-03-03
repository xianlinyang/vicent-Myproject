package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func copy() {
	//从s返回一个新的Reader,该reader含有大小,容量,可对其进行读写,偏移读写,重置等操作
	r1 := strings.NewReader("### this is read one new reader\n")
	r2 := strings.NewReader("### this is read two new reader\n")
	r3 := strings.NewReader("### this is read three new reader\n")

	//将字符串r copy到标准输出中打印出来
	io.Copy(os.Stdout, r1)

	//将字符串r copy到标准输出中打印出来,CopyN从src复制n个字节(或直到出错)到dst。
	io.CopyN(os.Stdout, r2, 100)

	//CopyBuffer与Copy是相同的，除了它逐步通过提供的缓冲区(如果需要一个)，而不是分配一个临时缓冲区。如果buf为nil，则分配一个;否则，如果它的长度为零，则CopyBuffer会恐慌。
	buf := make([]byte, 20)
	io.CopyBuffer(os.Stdout, r3, buf)
	log.Println(buf)
}

func read() {
	r1 := strings.NewReader("### this is write one  new reader\n")
	r2 := strings.NewReader("### this is write two  new reader\n")
	r3 := strings.NewReader("### this is write three  new reader\n")
	r4 := strings.NewReader("### this is write four  new reader\n")
	r5 := strings.NewReader("### this is write five  new reader\n")
	r6 := strings.NewReader("### this is write six  new reader\n")

	// ReadAll从r开始读取，直到出现错误或EOF，并返回它所读取的数据。
	ra, _ := io.ReadAll(r1)
	log.Println(string(ra))

	//可以一次读取多个,并拼接在一起
	re := io.MultiReader(r2, r3)
	io.Copy(os.Stdout, re)

	//将字符串读取到[]byte中,并设置至少读取长度,如果可读取的长度,则会报错,必须下面这个,r4字符串的长度决不能低于10,否则报错
	buf1 := make([]byte, 32)
	least, err := io.ReadAtLeast(r4, buf1, 10)
	if err != nil {
		log.Fatalln(err, least)
	}
	log.Println(string(buf1))

	// ReadFull从r读取len(buf)字节到buf。
	buf2 := make([]byte, 10)
	io.ReadFull(r5, buf2)
	log.Println(string(buf2))

	//LimitReader返回一个Reader，它从r中读取，但在n字节后以EOF结束。底层实现是一个*LimitedReader。
	lr := io.LimitReader(r6, 10)
	io.Copy(os.Stdout, lr)
}

func write() {
	//WriteString将字符串s的内容写入w, w接受一个字节片。如果w实现了一个WriteString方法，则直接调用它。否则，w.Write只调用一次。
	io.WriteString(os.Stdout, " this is write one string\n")

	//MultiWriter创建一个写入器，将其写入复制到所有提供的写入器。也就是说可以同时将一个数据写入到多个地方,下面将数据写入到两个buffer和一个标准输出中去
	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2, os.Stdout)
	r := strings.NewReader("this is write two string\n")
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
	log.Print(buf1.String())
	log.Print(buf2.String())
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//复制相关
	copy()

	//读取相关
	read()

	//写入相关
	write()

	// NewSectionReader返回一个从r读取的SectionReader
	//从offset off开始，在n字节后以EOF结束。
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)
	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}

	/**
	Pipe创建一个同步的内存管道。它可用于连接需要io的代码。带有io.Writer代码的阅读器。
	管道上的读和写是一一匹配的，除非多个读需要消耗一个写。也就是说，每个Write to the piperreader都会阻塞，直到它满足了来自piperreader的一个或多个read，这些read完全消耗了所写的数据。数据直接从写复制到相应的读(或读);没有内部缓冲。
	相互并行调用Read和Write或使用Close是安全的。对Read的并行调用和对Write的并行调用也是安全的:各个调用将按顺序门控。
	*/
	pr, pw := io.Pipe()
	go func() {
		//将字符串写入到writer中,写入后的数据可通过reader读取
		fmt.Fprint(pw, "this is i want write string\n")
		pw.Close()
	}()
	defer pr.Close()

	//将reader中读取的数据copy到标准输出中取
	if _, err := io.Copy(os.Stdout, pr); err != nil {
		log.Fatal(err)
	}

	//NopCloser返回一个ReadCloser函数，它带有一个无操作的Close方法来包装所提供的Reader r。
	rc := io.NopCloser(strings.NewReader("哈哈哈哈\n"))
	defer rc.Close()
	data, _ := io.ReadAll(rc)
	log.Println(string(data))
}
