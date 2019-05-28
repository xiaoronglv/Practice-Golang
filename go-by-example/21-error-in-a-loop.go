package main

import (
	"errors"
	"fmt"
)

func isEven(num int) (bool, error) {
	if num/2 == 0 {
		return true, nil
	} else {
		return false, errors.New("it's an odd!")
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, num := range nums {
		b, error := isEven(num)
		fmt.Println("Is it a even?", b)
		fmt.Println("the error is:", error)
	}
}
