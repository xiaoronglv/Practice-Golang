package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {
		fmt.Println("i is ", i)
	}

	j := 3

	for j <= 20 {
		j++
		fmt.Println("j is ", j)
	}

	for n := 0; n <= 30; n++ {
		if n%2 == 0 {
			fmt.Println("n is even?", n)
		}
	}
}
