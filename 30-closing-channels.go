package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		c1 <- "hello"
		close(c1)
	}()

	msg1, prs1 := <-c1
	msg2, prs2 := <-c1
	msg3, prs3 := <-c1

	fmt.Println(msg1)
	fmt.Println(prs1)

	fmt.Println(msg2)
	fmt.Println(prs2) // false

	fmt.Println(msg3) // no errors, amazing. so it means reading from closed channel will not
	// block.
	fmt.Println(prs3) // false

}
