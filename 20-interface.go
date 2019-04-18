package main

import (
	"fmt"
	"math"
	"reflect"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return (r.width + r.height) * 2
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g geometry) {
	t := reflect.TypeOf(g)
	fmt.Println("My type is ", t.Kind()) // the result is `struct`. Question: how to dig out the specific type?
	fmt.Println("My type name is ", t.Name())
	fmt.Println("I have _ fields", t.NumField())
	fmt.Println("area is ", g.area())
	fmt.Println("perim is", g.perim())
}

func main() {
	c := circle{radius: 333}
	r := rect{width: 32, height: 32}

	measure(c)

	measure(r)
}
