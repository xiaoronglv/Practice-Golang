package main

import (
	"fmt"
	"net/rpc"
	"rpc-demo/server"
)

func main() {
	args := &server.Args{1, 2}
	var reply int

	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	client.Call("Arith.Add", args, &reply)
	fmt.Println(reply)
}
