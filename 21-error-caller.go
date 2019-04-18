package main

import (
	"errors"
	"fmt"
)

func foo() {
	fmt.Println("I am foo")
	bar()
}

func bar() {
	fmt.Println(callers())
	fmt.Println("I am bar")
}

func main() {
	foo()
}
