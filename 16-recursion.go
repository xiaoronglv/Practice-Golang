package main

import (
	"fmt"
)

func decrease(num int) int {
	if num <= 0 {
		return num
	}

	num -= 1
	fmt.Println("the number is ", num)
	return decrease(num)
}

func main() {
	decrease(1000)
}
