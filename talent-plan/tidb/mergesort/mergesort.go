package main

import (
	"sort"
	"sync"
)

func merge(nums, tmp, left, right []int64) {
	j, k := 0, 0

	for i, _ := range nums {
		switch {
		case j >= len(left):
			tmp[i] = right[k]
			k++
		case k >= len(right):
			tmp[i] = left[j]
			j++
		case left[j] <= right[k]:
			tmp[i] = left[j]
			j++
		case left[j] > right[k]:
			tmp[i] = right[k]
			k++
		}
	}

	for i, _ := range tmp {
		nums[i] = tmp[i]
	}
}

func divide(nums []int64) ([]int64, []int64) {
	mid := len(nums) / 2
	return nums[0:mid], nums[mid:]
}

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	tmp := make([]int64, len(src))
	_mergeSort(src, tmp)
}

func _mergeSort(src, tmp []int64) {
	if len(src) < 100000 {
		sort.Slice(src, func(i, j int) bool {
			return src[i] < src[j]
		})
	} else {
		left, right := divide(src)
		lTmp, rTmp := divide(tmp)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			_mergeSort(left, lTmp)
		}()

		_mergeSort(right, rTmp)
		wg.Wait()
		merge(src, tmp, left, right)
	}
}
