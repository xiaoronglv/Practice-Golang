package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 4)

	c <- "hello"
	c <- "hello2"
	close(c)

	// 	go func() {
	// 		defer func() {
	// 			done <- true
	// 			fmt.Println("go routine is done")
	// 		}()
	// 		for s := range c {
	// 			fmt.Println(s)
	// 		}
	// 	}()

	for s := range c {
		fmt.Println(s)
	}

	fmt.Println("done")
}
