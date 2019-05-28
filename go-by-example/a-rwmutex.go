package main

import (
	"fmt"
	"sync"
)

type record struct {
	sync.RWMutex
	age  int
	name string
}

func main() {
	r := record{age: 33, name: "Ryan Lv"}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i <= 10000; i++ {
			r.Lock()
			r.age++
			fmt.Println("the age of r is ", r.age)
			r.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= 10000; i++ {
			r.RLock()
			fmt.Println("I am reading", r.age)
			r.RUnlock()
		}
		wg.Done()
	}()

	wg.Wait()
}
