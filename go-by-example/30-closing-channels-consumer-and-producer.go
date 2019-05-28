package main

import (
	"fmt"
	"time"
)

func worker(jobs <-chan string, done chan<- bool) {
	for {
		time.Sleep(time.Second * 2)
		j, more := <-jobs
		if more {
			fmt.Println(j)
		} else {
			done <- true
			return
		}
	}
}

func main() {
	jobs := make(chan string)
	done := make(chan bool)

	for i := 1; i <= 3; i++ {
		go worker(jobs, done)
	}

	for i := 1; i <= 10; i++ {
		jobs <- fmt.Sprintf("hello %d", i)
	}

	close(jobs)
	<-done
}
