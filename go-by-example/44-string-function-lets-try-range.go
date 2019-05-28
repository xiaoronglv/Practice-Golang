package main

import (
	"fmt"
)

func main() {
	s := "ab🔴d"

	for _, c := range s {
		fmt.Println(string(c))
	}
}
