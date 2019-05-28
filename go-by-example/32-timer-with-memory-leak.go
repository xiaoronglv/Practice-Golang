package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(100 * time.Second)

	go func() {
		defer func() {
			fmt.Println("I think this goroutine will hang there forever")
		}()

		fmt.Println("Do something before timer.")
		<-timer.C
		fmt.Println("Do something after timer.")
	}()

	timer.Stop()
	time.Sleep(3 * time.Second)

	fmt.Println("Time has already been stopped")
}
