package main

import (
	"fmt"
	_ "strings"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("true or false: %t\n", false)
	fmt.Printf("Print a decimal %d\n", 1031314143)
	fmt.Printf("Print a binary %b\n", 100)
	fmt.Printf("Print a hexadecimal %x\n", 100)
	fmt.Printf("Print a float %f\n", 100.1)
	fmt.Printf("what will happen if I print a integer in float format %f\n", 100)
	fmt.Printf("code point 233 stand for %c\n", 233)
	fmt.Printf("Print the hexadecimal of string %x\n", "é")
	fmt.Printf("Print the pointer %p\n", &p)
}
