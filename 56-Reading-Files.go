package main

import (
	_ "bufio"
	"fmt"
	_ "io"
	// "io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//dat, err := ioutil.ReadFile("/tmp/dat")
	// check(err)
	// fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)

	fmt.Println(n1)
	fmt.Println(string(b1))
}
