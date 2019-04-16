package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := rand.Intn(100)

	if a%2 == 0 {
		fmt.Println("a is an even", a)
	} else {
		fmt.Println("a is an odd", a)
	}

	if a%10 == 0 {
		fmt.Println("a is divisible by 10", a)
	} else {
		fmt.Println("a is not a decemial number", a)
	}

	rand.Seed(time.Now().UnixNano())
	money := rand.Intn(100)

	if money > 50 {
		fmt.Println("this man is a rich man, since he has lots of money. ", money)
	} else if money > 20 {
		fmt.Println("this man is a middle class, since he has 20-50 yuan.", money)
	} else {
		fmt.Println("this man is poor, since he has few money.", money)
	}
}
