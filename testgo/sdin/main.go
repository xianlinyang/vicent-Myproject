package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//stdReader := bufio.NewReader()
	go func(){
		stdReader := bufio.NewReader(os.Stdin)
		endData, _ := stdReader.ReadString('\n')
		fmt.Println(endData)
	}()
	select {}
}


