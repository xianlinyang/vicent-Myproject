package rpc

import "net/rpc"

func main() {
	rpc.Register()
	rpc.HandleHTTP()
}
