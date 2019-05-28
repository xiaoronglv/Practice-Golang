package main

import (
	"fmt"
)

type article struct {
	title     string
	body      string
	author_id int
}

func main() {
	a := article{title: "Who is Ryan?", body: "Ryan is a developer", author_id: 3}
	fmt.Println(a)

	b := article{title: "Who is angela"}
	fmt.Println("print a ommited field", b.author_id)

	bp := &b
	bp.author_id = 3
	fmt.Println("set new value through pointer", bp.author_id)
	b.author_id = 4
	fmt.Println("set new value through struct", b.author_id)
	// so I can draw the conclusion that we can set value through struct or pointer, they both work.
}
