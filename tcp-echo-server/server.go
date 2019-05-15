package main

import (
	"bufio"
	"fmt"
	"net"
)

var reqs = make(chan string)

func handleRequest(conn net.Conn) {
	io := bufio.NewReader(conn)
	for {
		req, err := io.ReadString('\n')
		if err != nil {
			break
		}

		reqs << // immediately write to all channels of client
			conn.Write([]byte(req))
		fmt.Println(req)
	}
}

type readOp struct {
	resp chan chan string
}

type addOp struct {
	value chan string
	resp  chan bool
}

// note:
// https://blog.golang.org/pipelines
type Broadcast struct {
	readOps   chan *readOp
	addOps    chan *addOp
	publisher chan string
	// how to safely add/remove subscribers?
	// How about stateful goroutine?
	subscribers []chan string
}

func (b *Broadcast) listen() {
	go func() {
		for {
			select {
			case read := <-b.readOps:
				read.resp <- b.subscribers
			case add := <-b.addOps:
				b.subscribers = append(b.subscribers, add.value)
				add.resp <- true
			}
		}
	}()

	for {
		select {
		case v1 := <-b.publisher:
			for _, c := range b.allSubscribers {
				select {
				case c <- v1:
				default:
				}
			}

		}
	}
}

func (b *Broadcast) allSubscribers() []chan string {
	read := &readOp{resp: make(chan []chan string)}
	reads <- read
	<-read.resp
}

func (b *Broadcast) addSubscribers(chan string) {
	add := &addOp{resp: make(chan bool)}
	adds <- add
	<-add.resp
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	go func() {

		for {
			conn, err := ln.Accept()
			defer conn.Close()

			if err != nil {
				panic(err)
			}

			go handleRequest(conn)
		}
	}()

	select {}
}
