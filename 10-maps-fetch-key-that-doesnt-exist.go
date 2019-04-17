package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)

	value, prs := m["key_doesnt_exist"]

	fmt.Println("the value is ", value)
	fmt.Println("is it present?", prs)
}
