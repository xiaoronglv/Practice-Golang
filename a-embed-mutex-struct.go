package main

import (
	"fmt"
	"sync"
	"time"
)

type master struct {
	sync.Mutex
	A int
}

func main() {
	mr := &master{A: 1}

	go func() {
		time.Sleep(2 * time.Second)
		mr.Lock()
		mr.A++
		fmt.Println("A is updated in goroutine d%", mr.A)
		mr.Unlock()
	}()

	mr.Lock()
	time.Sleep(30 * time.Second)
	fmt.Println("A is updated in main")
	mr.A++
	mr.Unlock()
}
