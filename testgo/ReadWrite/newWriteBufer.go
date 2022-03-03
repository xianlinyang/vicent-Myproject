package main

import (
	"bufio"
	"fmt"
)

type Writer1 int

func (*Writer1) Write(p []byte) (n int, err error) {
	fmt.Println("fff", len(p))
	return len(p), nil
}
func main() {
	fmt.Println("Unbuffered I/O")
	w := new(Writer1)
	w.Write([]byte{'a'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})
	w.Write([]byte{'d'}) //每一次都会调用write
	fmt.Println("Buffered I/O")
	bw := bufio.NewWriterSize(w, 3) //这里写三次才会调用一次
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})
	err := bw.Flush()
	if err != nil {
		panic(err)
	}
}
