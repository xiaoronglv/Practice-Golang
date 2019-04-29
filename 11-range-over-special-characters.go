package main

import (
	"fmt"
)

func main() {
	s3 := "haäé"

	for _, c := range s3 {
		fmt.Println(string(c))
	}
}
