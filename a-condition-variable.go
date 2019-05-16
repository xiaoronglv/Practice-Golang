package main

import (
	"fmt"
	"sync"
)

func main() {
	l := sync.Mutex{}
	c := sync.NewCond(&l)
	var wg sync.WaitGroup

	m := make(map[int]int)

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			l.Lock()
			defer l.Unlock()
			defer wg.Done()

			if _, ok := m[i]; ok {
				return
			} else {
				c.Wait()
				fmt.Println("the broadcast message has been well received")
				m[i] = i * i
			}
		}()

	}

	l.Lock()
	m[0] = 0
	c.Broadcast()
	fmt.Println("I have already broadcast this message")
	l.Unlock()

	wg.Wait()
}
