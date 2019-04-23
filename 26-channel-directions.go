package main

import (
	"fmt"
)

// expected a compile error
// func readMsg(pings chan<- string) string {
// 	msg := <-pings
// }

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	fmt.Println("I am echoing")
	fmt.Println(msg)
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hello")
	pong(pings, pongs)

}
