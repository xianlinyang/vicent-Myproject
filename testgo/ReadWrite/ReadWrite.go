package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {

	rfile, _ := ioutil.ReadFile("/home/xianlin_yang/Desktop/nodetest/fileTest/kad/2.jpg")

	s1 := bytes.NewBuffer(nil)
	n, _ := s1.Write(rfile)
	fmt.Println(n)
	fmt.Println(s1.Len())
	k1 := make([]byte, s1.Len())

	_, _ = s1.Read(k1)
	_ = ioutil.WriteFile("/home/xianlin_yang/Desktop/nodetest/fileTest/kad/4.jpg", k1, 0644)

}
