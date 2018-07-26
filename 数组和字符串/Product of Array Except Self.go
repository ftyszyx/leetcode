package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	arrnum := len(nums)
	var left = make([]int, arrnum)
	for i := 0; i < arrnum; i++ {
		if i == 0 {
			left[i] = 1
		} else {
			left[i] = nums[i-1] * left[i-1]
		}
	}

	var right = make([]int, arrnum)
	for i := arrnum - 1; i >= 0; i-- {
		if i == arrnum-1 {
			right[i] = 1
		} else {
			right[i] = right[i+1] * nums[i+1]
		}
	}
	var res = make([]int, arrnum)
	for i := 0; i < arrnum; i++ {
		res[i] = left[i] * right[i]
	}
	return res
}

func main() {
	input := []int{1, 2, 3, 4}
	println(fmt.Sprintf("%#v", productExceptSelf(input)))
}
