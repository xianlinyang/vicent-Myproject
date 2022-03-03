package main

import "fmt"

func gofunc(ch chan bool, i int) {
	fmt.Println("ss", i)
	ch <- true
}
func main() {
	ch := make(chan bool)

	for i := 0; i < 10; i++ {
		go gofunc(ch, i)
		<-ch
		fmt.Println("sss")
	}

}
