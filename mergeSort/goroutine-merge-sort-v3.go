package main

import (
	"fmt"
)

func merge(nums, left []int64, right []int64) {
	j, k := 0, 0

	for i, _ := range nums {
		switch {
		case j >= len(left):
			nums[i] = right[k]
			k++
		case k >= len(right):
			nums[i] = left[j]
			j++
		case left[j] < right[k]:
			nums[i] = right[k]
			k++
		case left[j] > right[k]:
			nums[i] = left[j]
			j++
		case left[j] == right[k]:
			nums[i] = left[j]
			j++
		}
	}
}

func divide(nums []int64) ([]int64, []int64) {
	mid := len(nums) / 2
	return nums[0:mid], nums[mid:]
}

func sort(nums []int64) {
	if len(nums) <= 1 {
		return
	}

	left, right := divide(nums)

	Done := make(chan bool)
	go func() {
		sort(left)
		Done <- true
	}()

	sort(right)
	<-Done

	merge(nums, left, right)
}

func main() {
	nums := []int64{3, 9, 7, 1, 3, 9, 8, 10, 23, 22, 97, 34, 5, 2, 91}
	sort(nums)
	fmt.Println(nums)
}
