package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Intn(5)

	fmt.Println(number)

	switch number {
	case 0:
		fmt.Println("I am zero")
	case 1:
		fmt.Println("I am one")
	case 2:
		fmt.Println("I am two")
	case 3:
		fmt.Println("I am three")
	case 4:
		fmt.Println("I am four")
	}
}
