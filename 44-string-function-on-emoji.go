package main

import (
	"fmt"
	//	"strings"
)

var p = fmt.Println

func main() {
	s := "ab🔴d"
	p("Char 0 should be a:", string(s[0])) // works fine
	p("Char 1 should be b:", string(s[1])) // works fine
	p("Char 2 is ?:", string(s[2]))        // the result is not 🔴
	p("Char 3 is ?:", string(s[3]))        // the result is not d
}
