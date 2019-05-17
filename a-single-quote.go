package main

import (
	"fmt"
)

func main() {
	a := `I am rune?`

	for _, r := range a {
		fmt.Println(r)
	}
}
