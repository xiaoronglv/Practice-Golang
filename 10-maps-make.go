package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["foo"] = "I am foo"
	m["bar"] = "I am bar"

	fmt.Println(m)
}
