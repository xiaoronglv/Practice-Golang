package main

import (
	"fmt"
)

func plus(a, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func main() {
	a := 1
	b := 2
	c := 3

	fmt.Println("a+b+c is ", plusplus(a, b, c))
	fmt.Println("a+b is ", plus(a, b))
}
