package main

import (
	"fmt"
)

func sum(nums ...int) int {
	sum := 0
	for num := range nums {
		sum += num
	}

	return sum
}

func main() {
	fmt.Println("sum 1, 2, 3, 4, 5 =", sum(1, 2, 3, 4, 5))

	nums := []int{1, 2, 3, 4, 5}
	total := sum(nums...)

	fmt.Println("total is equal to", total)
}
