// there is a bug in this file.
package main

import (
	"fmt"
)

func worker(jobs <-chan int, results chan<- string) {
	for {
		job, prs := <-jobs
		if prs {
			results <- fmt.Sprint("job is done %d", job)
		} else {
			close(results)
		}
	}
}

func main() {
	jobs := make(chan int, 3)
	results := make(chan string, 3)

	go func() {

		for i := 1; i <= 100; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	for i := 1; i <= 5; i++ {
		go worker(jobs, results)
	}

	for r := range results {
		fmt.Println(r)
	}
}
