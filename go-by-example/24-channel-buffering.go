package main

import (
	"fmt"
)

func main() {
	msgs := make(chan string, 3)

	go func() {
		msgs <- "I"
		msgs <- "am"
		msgs <- "Ryan"
	}()

	msg1 := <-msgs
	msg2 := <-msgs
	msg3 := <-msgs

	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
}
