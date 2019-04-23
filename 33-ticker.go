package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	for t := range ticker.C {
		fmt.Println("tick", t)
	}
}
