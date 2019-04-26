package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(100 * time.Second)
	stop := make(chan bool)

	go func() {
		defer func() {
			fmt.Println("tear down")
		}()

		select {
		case <-timer.C:
			fmt.Println("Do something after timer.")
		case <-stop:
			fmt.Println("timer is stopped")
		}
	}()

	timer.Stop()
	stop <- true

	time.Sleep(3 * time.Second)

	fmt.Println("Time has already been stopped")
}
