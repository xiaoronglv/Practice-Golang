package main

import (
	"bufio"
	"fmt"
	"net"
)

var reqs = make(chan string)
var newSubs = make(chan chan string)
var subs = make([]chan string, 5) // TODO: how to write a slice of channels
// var unsubs = make([]chan string, 5)

func handleRequest(conn net.Conn) {
	io := bufio.NewReader(conn)
	sub := make(chan string, 100)
	newSubs <- sub

	go func() {
		for {
			req, err := io.ReadString('\n')
			if err != nil {
				break
			}
			reqs <- req
		}
	}()

	for {
		resp := <-sub
		fmt.Println(resp)
		conn.Write([]byte(resp))
	}

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

	go func() {
		for {
			select {
			case sub := <-newSubs:
				subs = append(subs, sub)
			case req := <-reqs:
				fmt.Println("range channel before")
				for _, c := range subs {
					select {
					case c <- req:
					default:
						// "do nothing, just change this write to non-blocking"
						fmt.Println("buffer is full, discard the message")

					}
				}
				//	case unsubC := <-unsubs:
				//		var i int
				//		for i, c := range subs {
				//			if c == unsubC {
				//				return i
				//			}
				//			}
				//		subs[i].Close()
				//		subs = append(subs[:i], subs[i+1:]...)
			}
		}
	}()

	select {}
}
