package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() {
		messages <- "hello world"
		messages <- "hello world again"
	}()

	msg := <-messages
	fmt.Println(msg)

	msg2 := <-messages
	fmt.Println(msg2)

}
