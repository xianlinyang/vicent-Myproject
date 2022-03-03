package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

import (
	"context"
)

type test01String string
type test02String string
const (
	GOLABLE_KEY = "test123"
)



func test01(ctx context.Context) {
	// context实例
	ctx01 := context.WithValue(ctx, test01String("test01"), "hello1")
	test02(ctx01)
}

func test02(ctx context.Context) {
	// context实例
	ctx02 := context.WithValue(ctx, test02String("test02"), "hello2")
	test03(ctx02)
}

func test03(ctx context.Context) {
	fmt.Println(ctx.Value(test01String("test01")))
	fmt.Println(ctx.Value(test02String("test02")))
}




func controlAllConrrent() {

	logg := log.New(os.Stdout, "", log.Ltime)
	handleSome()
	logg.Println("over ")
}

//父协程
func handleSome() {
	//ctx, cancelFunc := context.WithCancel(context.Background())
	//ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))

	ctx = context.WithValue(ctx, GOLABLE_KEY, "1234")

	go workerFunc(ctx, "str1")
	go workerFunc(ctx, "str2")
	time.Sleep(time.Second * 3)
	cancelFunc()
}

//子协程
func workerFunc(ctx context.Context, showStr string) {
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			log.Println("done")
			fmt.Println(ctx.Err())
			return
		default:
			val123 := ctx.Value(GOLABLE_KEY).(string)
			log.Println(val123, showStr)
		}
	}
}

func main() {
	//// context实例
	//ctx := context.Background()
	//test01(ctx)

	//父context控制子context
	//by := make([]byte,1)
	//by[0]  ^= 0x1 << 3
	//fmt.Printf("%08b ",by)
	////controlAllConrrent()
	rand.Seed(time.Now().Unix())
	fmt.Println(time.Now().Unix())
	fmt.Println(rand.Int63())

	fmt.Println(8 % 3)
}