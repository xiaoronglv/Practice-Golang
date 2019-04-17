package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)

	// set
	s[1], s[2], s[4] = 2314, 214124, 3141

	// append
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	// pop
	new_s := s[0 : len(s)-1]
	fmt.Println(new_s)

	// unshift1
	nns := make([]int, 1)
	nns[0] = 123414112414
	nns = append(nns, s...)

	fmt.Println(nns)

	// unshift2
	nns = concat([]int{4321}, nns)
	fmt.Println(nns)
}
