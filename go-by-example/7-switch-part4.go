package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("It's ", time.Now().Weekday())
	fmt.Println(time.Saturday)

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	}
}
