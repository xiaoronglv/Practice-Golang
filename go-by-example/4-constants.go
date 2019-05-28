package main

import (
	"fmt"
	_ "math"
	"reflect"
)

const s string = "I am a constant string"

func main() {
	fmt.Println(s)

	const n = 5000000000
	const d = 3e20 / n

	t := reflect.TypeOf(d)

	fmt.Println("t is a ?", t)
	fmt.Println("d is float? the answer is ", d)
	fmt.Println("I am trying to convert d to int", int64(d))
}
