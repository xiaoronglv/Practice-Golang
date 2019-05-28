package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"I", "am", "Ryan", "Lv"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{1, 9, 8, 6, 0, 6, 0, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	sorted := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", sorted)
}
