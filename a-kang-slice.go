package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3)
	// append(s, 1) // it doesn't work
	// append(s, 2)
	// append(s, 3)
	// append(s, 4)
	// append(s, 5)
	// append(s, 6)
	// append(s, 7)
	// append(s, 8)
	// append(s, 9)
	// append(s, 10)
	// append(s, 11)

	// the right way to append a slice
	s = append(s, 11)
	fmt.Println(s) // Kang is right.
}
