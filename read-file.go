package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("/tmp/go-by-example-data/foo.txt")
	fmt.Println(string(data))
}
