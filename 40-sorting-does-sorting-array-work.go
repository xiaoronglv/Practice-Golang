package main

import (
	"fmt"
	"sort"
)

func main() {
	names := [4]string{"Ryan", "Sam", "Lv", "Qiu"} // it doesn't work
	sort.Strings(names)
	fmt.Println(names)
}
