package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string)

	// read
	select {
	case v1 := <-c1:
		fmt.Println(v1)
	default:
		fmt.Println("print default")
	}

	// send
	select {
	case c1 <- "hello":
		fmt.Println("hello is enqueued")
	default:
		fmt.Println("nothing is enqueued")
	}
}
