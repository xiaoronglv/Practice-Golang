package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"aaaaa", "a", "aaa", "aaaaaaa"}

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	fmt.Println("Sorted Slices:", strs)
}
