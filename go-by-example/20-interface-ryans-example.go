package main

import (
	"fmt"
)

type Eatable interface {
	/* TODO: add methods */
	eat()
}

type dog struct {
	name   string
	color  string
	age    int
	gender int
}

func (d dog) eat() {
	fmt.Println("I am dog, and I can eat something")
}

type baby struct {
	name string
}

func (baby baby) eat() {
	// question: what's the return value when there is no return?
	fmt.Println("I am a baby and I can eat something")
}

func feed(something Eatable) {
	something.eat()
}

func main() {
	dog := dog{name: "Golden retriever"}
	feed(dog)
	baby := baby{name: "Bush"}
	feed(baby)
	a := func() int {
		return 3
	}()
	fmt.Println(a)
}
