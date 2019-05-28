package main

import (
	"fmt"
)

func main() {
	// slice

	nums := []int{1, 2, 3, 4, 5, 1000, 10000, 10000000}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("the sum is ", sum)

	// array

	array := [4]int{1, 9, 8, 7}
	sum2 := 0

	for _, num := range array {
		sum2 += num
	}

	fmt.Println("the sum2 is ", sum2)

}
