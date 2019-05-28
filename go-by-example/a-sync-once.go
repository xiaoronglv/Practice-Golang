package main

import (
	"fmt"
	"sync"
)

type member struct {
	s    sync.Once
	name string
	age  int
}

func (m *member) setup() {
	m.s.Do(func() {
		m.age = 33
		m.name = "Ryan Lv"
		fmt.Println("I only run once")
	})
}

func main() {
	m := member{}
	for i := 0; i <= 100; i++ {
		m.setup()
	}
}
