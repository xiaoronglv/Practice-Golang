package main

import (
	"fmt"
	_ "time"
)

func puts(s string) {
	fmt.Println(s)
}

func main() {
	go puts("1")
	go puts("2")
	go puts("3")
	go puts("4")
	go puts("5")
	go puts("6")
	go puts("7")
	puts("0")
}
