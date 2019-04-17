package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"foo": "value of foo", "bar": "value of bar"}
	delete(m, "foo")

	fmt.Println(m)
}
