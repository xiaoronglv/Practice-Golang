package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(4 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("do something in goroutine")
	}()

	stopped := timer.Stop()
	if stopped {
		fmt.Println("timer is stopped")
	}
	time.Sleep(10 * time.Second)
}
