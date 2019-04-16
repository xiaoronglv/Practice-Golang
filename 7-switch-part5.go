package main

import (
	"fmt"
)

func main() {
	whoiam := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("i am a int", i)
		case bool:
			fmt.Println("i am a boolean", i)
		}
	}

	whoiam(true)
	whoiam(123414)
}
