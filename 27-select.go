package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			c1 <- "hello from loop1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			c2 <- "hello from loop2"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			c3 <- "hello from loop3"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg3 := <-c3:
			fmt.Println("received", msg3)
		}
	}
}
