package main

import (
	"encoding/json"
	"os"
)

func main() {
	m := make(map[string]string)
	m["Ryan"] = "I"
	m["Angela"] = "Love"
	m["Solomon"] = "You"

	f, _ := ioutil.WriteFile("/tmp/go-by-example-data/foo.txt")
}
