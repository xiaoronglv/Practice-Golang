package main

func merge(left []int64, right []int64) []int64 {

}

func divide(nums []int64) ([]int64, []int64) {
	mid := len(nums) / 2
	return nums[0:mid], nums[mid:]
}

func sort(nums []int64) []int64 {
	if len(num) <= 1 {
		return nums
	}

	left, right := divide(nums)
	done := make(chan bool)

	go func() {
		left = sort(left)
		right = sort(right)
		done <- true
	}()

	<-done
	return merge(left, right)
}

func main() {

}
