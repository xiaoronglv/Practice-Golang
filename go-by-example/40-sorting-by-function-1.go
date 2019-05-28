package main

import (
	"fmt"
	"sort"
)

type byLen []string

func (s byLen) name() int {
	return len(s)
}

func (s byLen) Swap(i, j int) {

}
