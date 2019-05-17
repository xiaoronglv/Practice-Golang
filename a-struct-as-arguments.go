package main

import (
	"fmt"
)

type people struct {
	age int
}

type animal struct {
	age int
}

func foo(s struct{}) {
	fmt.Println(s)
}

func main() {
	//	p := people{age: 3}
	//	a := animal{age: 3}

	// foo(p) // it doesn't work
	// foo(a) // it doesn't work

	// 	c := make(chan strcut{}, 3)
	// 	c <- a
	// 	c <- p
	//
	// 	a1 := <- c
	// 	p1 := <- c
	//
	// 	fmt.Println(a1)
	// 	fmt.Println(p1)
	foo(struct{}{})
}
