package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Result 1"
	}()

	// Question:
	// how to handle IO timeout?
	// how to handle Non-IO timeout?
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("time out")
	}
}
