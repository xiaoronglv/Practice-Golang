package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 2, 39199239, 12319, 31324, 1232, 9999, 221, 30}

	for _, num := range s {
		fmt.Println(num)
		fmt.Printf("Addr: %p\n", &num)
	}

	c := make(chan int, 20)

	for i := 0; i <= 19; i++ {
		c <- i
	}

	close(c)

	for n := range c {
		fmt.Println(n)
		fmt.Printf("Addr: %p\n", &n)
	}

}
