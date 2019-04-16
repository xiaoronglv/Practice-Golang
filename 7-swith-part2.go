package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Intn(4)

	switch {
	case number%2 == 0:
		fmt.Println("number is an even", number)
	case number%2 == 1:
		fmt.Println("number is an odd", number)
	}
}
