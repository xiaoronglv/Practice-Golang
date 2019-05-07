package main

import "fmt"

func divide(nums []int64) ([]int64, []int64) {
	mid := len(nums) / 2
	return nums[0:mid], nums[mid:]
}

func merge(left []int64, right []int64) []int64 {
	array := make([]int64, len(left)+len(right))
	l, r := 0, 0

	for i := 0; i < len(array); i++ {
		if l >= len(left) {
			array[i] = right[r]
			r++
			continue
		}

		if r >= len(right) {
			array[i] = left[l]
			l++
			continue
		}

		if left[l] < right[r] {
			array[i] = right[r]
			r++
		} else {
			array[i] = left[l]
			l++
		}
	}

	return array
}

func sort(nums []int64) []int64 {
	if len(nums) <= 1 {
		return nums
	}

	left, right := divide(nums)
	left = sort(left)
	right = sort(right)
	return merge(left, right)
}

func main() {
	nums := []int64{1, 2, 1111, 939393, 38, 98, 321, 3, 7, 11, 4, 0, 9, 1000, 21}
	fmt.Println(sort(nums))
}
