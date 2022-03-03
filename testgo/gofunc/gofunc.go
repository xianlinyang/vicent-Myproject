package main

import "fmt"

func main(){
	k := FuncTest()
	fmt.Println(k(1))
}
func FuncTest()(Rfunc func(a int) int){
	var k int
	k = 10
	Rfunc= func(a int) int{
		return a + k
	}
	return Rfunc
}

