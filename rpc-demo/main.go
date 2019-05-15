package main

import (
	"net"
	"net/http"
	"net/rpc"
	. "rpc-demo/server"
)

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, _ := net.Listen("tcp", ":1234")

	go http.Serve(l, nil)
	select {}
}
