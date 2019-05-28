package main

import (
	"fmt"
)

type rect struct {
	length int
	height int
}

// quesiton: what's the difference between the following 2 methods?
// pointer maybe used to manipulate the state of this struct.
func (rect rect) area() int {
	return rect.length * rect.height
}

func (r *rect) area2() int {
	return r.length * r.height
}

func main() {
	shape := rect{length: 16, height: 4}
	fmt.Println("the area of this shape is ?", shape.area())

	ptr := &rect{length: 16, height: 4}
	fmt.Println(" the area of this shape is calculated through pointer", ptr.area2())

}
