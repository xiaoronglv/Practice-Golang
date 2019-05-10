package main

import (
	"fmt"
)

func merge(left []int64, right []int64) []int64 {
	array := make([]int64, len(left)+len(right))
	j, k := 0, 0

	for i, _ := range array {
		switch {
		case j >= len(left):
			array[i] = right[k]
			k++
		case k >= len(right):
			array[i] = left[j]
			j++
		case left[j] < right[k]:
			array[i] = right[k]
			k++
		case left[j] > right[k]:
			array[i] = left[j]
			j++
		case left[j] == right[k]:
			array[i] = left[j]
			j++
		}
	}

	return array
}

func divide(nums []int64) ([]int64, []int64) {
	mid := len(nums) / 2
	return nums[0:mid], nums[mid:]
}

func sort(nums []int64) []int64 {
	if len(nums) <= 1 {
		return nums
	}

	left, right := divide(nums)

	lDone := make(chan bool)
	rDone := make(chan bool)

	go func() {
		left = sort(left)
		lDone <- true
	}()

	go func() {
		right = sort(right)
		rDone <- true
	}()

	<-lDone
	<-rDone

	return merge(left, right)
}

func main() {
	nums := []int64{3, 9, 7, 1, 3, 9, 8, 10, 23, 22, 97, 34, 5, 2, 91}
	fmt.Println(sort(nums))
}
