package main

import (
	"fmt"
	"sync"
)

type waitCounter struct {
	counter int
	mutex   sync.Mutex
	allDone chan bool
}

func (w *waitCounter) Add(counter int) {
	w.counter = counter
	w.mutex = sync.Mutex{}
	w.allDone = make(chan bool, 1)
}

func (w *waitCounter) Wait() {
	<-w.allDone
}

func (w *waitCounter) Done() {
	w.mutex.Lock()
	w.counter = w.counter - 1
	w.mutex.Unlock()

	if w.counter == 0 {
		w.allDone <- true
		close(w.allDone)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func mergeSort(nums []int64) {
	low := 0
	high := len(nums) - 1
	temp := make([]int64, len(nums))

	for step := 1; step <= high-low; step = 2 * step {
		//		counterWait := &counterWait{}
		//		counterwait.Add(len(nums)/step + 1)

		for i := low; i < high; i += 2 * step {
			from := i
			mid := i + step - 1
			to := min((i + 2*step - 1), high)

			merge(nums, temp, from, mid, to)
		}

		//		counterWait.Wait()
	}
}

func merge(nums []int64, temp []int64, low int, mid int, high int) {
	i := low     // index for temp
	j := low     // index of left halve
	k := mid + 1 // index of right halve

	switch {
	case j > mid:
		temp[i] = nums[k]
		k++
		i++
	case k > high:
		temp[i] = nums[j]
		j++
		i++
	case nums[j] < nums[k]:
		temp[i] = nums[j]
		j++
		i++
	case nums[j] > nums[k]:
		temp[i] = nums[k]
		k++
		i++
	}

	for i := low; i <= high; i++ {
		nums[i] = temp[i]
	}
}

func main() {
	nums := []int64{1, 2, 3, 23, 21, 31, 1231, 9, 2, 0, 92, 123}
	mergeSort(nums)
	fmt.Println(nums)
}
