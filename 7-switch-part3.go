package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Intn(6)

	switch i {
	case 1, 2, 3, 0:
		fmt.Println("it's a small number", i)
	case 4, 5, 6:
		fmt.Println("it's a large number", i)
	}
}
