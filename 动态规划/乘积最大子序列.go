package main

import (
	"fmt"
)

func main() {

	var arr = []int{2, 3, -2, 4}
	fmt.Println(fmt.Sprintf("%d", maxProduct(arr)))
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxProduct(nums []int) int {
	var negmax []int //每个点
	var posmax []int //每个点的正数

	var maxres = 0

	for index, item := range nums {

		if index == 0 {
			negmax = append(negmax, item)
			posmax = append(posmax, item)
			maxres = item
		} else {
			posmax = append(posmax, max(item, max(item*posmax[index-1], item*negmax[index-1])))
			negmax = append(negmax, min(item, min(item*posmax[index-1], item*negmax[index-1])))
			maxres = max(maxres, posmax[index])
		}

	}
	return maxres
}
