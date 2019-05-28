package main

import (
	"fmt"
)

func main() {
	var a [5]int

	a[0], a[1], a[4] = 1, 2, 3

	fmt.Println(a)
}
