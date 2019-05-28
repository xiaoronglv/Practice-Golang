package main

import (
	"fmt"
)

func foo(f func(string)) {
	f("hello")
}

func main() {
	f := func(s string) {
		fmt.Println(s)
	}

	foo(f)
}
