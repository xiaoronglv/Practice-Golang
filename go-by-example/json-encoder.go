package main

import (
	"encoding/json"
	"os"
)

func main() {
	a := "Ryan is eating"
	b := false
	c := 3
	d := map[string]int{"apple": 3, "lette": 4}
	e := "d"

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(a)
	enc.Encode(b)
	enc.Encode(c)
	enc.Encode(d)
	enc.Encode(e)
}
