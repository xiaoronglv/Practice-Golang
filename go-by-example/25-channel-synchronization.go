package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("I am doing something")
	time.Sleep(time.Second)
	fmt.Println("I am done in worker")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
