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
		time.Sleep(1 * time.Second)
		for i := 0; i < 100; i++ {
			go func() {
				mr.A++
				fmt.Println("I am updating when main is sleeping")
			}()
		}

	}()

	mr.Lock()
	mr.A++
	time.Sleep(10 * time.Second)
	mr.Unlock()
	fmt.Println(mr.A)

}
