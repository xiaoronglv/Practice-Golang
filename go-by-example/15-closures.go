package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := intSeq()

	fmt.Println("next integer is ", nextInt())
	fmt.Println("next integer is ", nextInt())
	fmt.Println("next integer is ", nextInt())
	fmt.Println("next integer is ", nextInt())
	fmt.Println("next integer is ", nextInt())

	newInts := intSeq()
	fmt.Println("newInts integer is ", newInts())
	fmt.Println("newInts integer is ", newInts())
	fmt.Println("newInts integer is ", newInts())
}
