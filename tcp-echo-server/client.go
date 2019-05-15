package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100000000000; i++ {
		s := fmt.Sprintf("%d\n", i)
		fmt.Println(s)
		fmt.Fprintf(conn, s)
	}
}
