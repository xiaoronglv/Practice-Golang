package main

import (
	"fmt"
	"math/rand"
	"time"
)

var nameList = []string{
	"Ryan",
	"Vincent",
	"Felix",
	"Ricky",
	"Tao",
	"Shaun",
	"Hazel",
}

type readOp struct {
	key  string
	resp chan int
}

type writeOp struct {
	key   string
	value int
	resp  chan int
}

func main() {
	readC := make(chan *readOp)
	writeC := make(chan *writeOp)
	rand.Seed(time.Now().UnixNano())

	go func() {
		state := make(map[string]int)
		select {
		case readOp := <-readC:
			readOp.resp <- state[readOp.key]
		case writeOp := <-writeC:
			state.value = writeOp.value
			writeOp.resp <- true
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := &readOp{
					key:  nameList[rand.Int(len(nameList)-1)],
					resp: make(chan bool),
				}

				readC <- read
				response := <-read.resp
				fmt.Println("Response: ", response)
				time.Sleep(time.Millisecond)
			}

		}()

	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				write := &writeOp{
					key:   nameList[rand.Int(len(nameList)-1)],
					value: rand.Int(10000),
					resp:  make(chan bool),
				}

				writeC <- write
				response := <-write.resp

				fmt.Println("write is done:", response)
				time.Sleep(time.Millisecond)
			}
		}()

	}

}
