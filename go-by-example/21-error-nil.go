package main

import (
	"fmt"
)

func foo() (int, error) {
	return 1, nil
}

func main() {
	a, b := foo()

	fmt.Println("a is ", a) // it seems it doesn't work
	fmt.Println("b is ", b) // it seems it doesn't work
	// nil is not a type.
	// interesting
}
