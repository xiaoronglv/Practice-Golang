package main

import (
	"fmt"
)

// scenario of point
type student struct {
	name   string
	age    int
	gender int
	status string
}

func (s *student) graduate() {
	s.status = "graduated"
}

func main() {
	p := &student{"Ryan Lv", 32, 1, "in-progress"}
	p.graduate()

	fmt.Println("Print out Ryan Lv", p.status)
}
