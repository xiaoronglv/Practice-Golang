package main

import (
	"fmt"
	"sync"
)

type waitCounter struct {
	counter int
	mutex   sync.Mutex
	allDone chan bool
}

func (w *waitCounter) Add(counter int) {
	w.counter = counter
	w.mutex = sync.Mutex{}
	w.allDone = make(chan bool, 1)
}

func (w *waitCounter) Wait() {
	<-w.allDone
}

func (w *waitCounter) Done() {
	w.mutex.Lock()
	w.counter = w.counter - 1
	w.mutex.Unlock()

	if w.counter == 0 {
		w.allDone <- true
		close(w.allDone)
	}
}

func main() {
	w := &waitCounter{}
	w.Add(4)
	w.Done()
	w.Done()
	w.Done()
	w.Wait()
	fmt.Println(1)
}
