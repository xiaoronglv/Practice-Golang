package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	go func() {
		for i := 0; i <= 10; i++ {
			c <- fmt.Sprintf("hello%d", i)
		}

		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}

}
