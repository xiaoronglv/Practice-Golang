package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(10 * time.Second)

	go func() {
		ticker := time.NewTicker(time.Second)
		for t := range ticker.C {
			fmt.Println("tick, tick, tick", t)
		}
	}()

	<-timer.C
	fmt.Println("finished")
}
