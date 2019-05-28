package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				time.Sleep(time.Nanosecond)
			}
		}()
	}

	for i := 0; i < 50; i++ {
		go func() {
			for {
				fmt.Println(ops)
				time.Sleep(time.Nanosecond)
			}
		}()
	}
	time.Sleep(100 * time.Second)

	opsFinal := atomic.LoadUint64(&ops) // Can I replace it?
	fmt.Println("ops:", opsFinal)
}
