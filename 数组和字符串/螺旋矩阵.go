package main

import (
	"fmt"
)

func loopmatrix(matrix [][]int, level int, res []int, index int) {
	row := len(matrix)
	col := len(matrix[0])
	total := row * col
	if index >= total {
		return
	}
	rowstart := level
	rowend := row - level - 1

	colstart := level
	colend := col - level - 1

	// println(fmt.Sprintf("row start:%d rowend:%d colstart:%d colend:%d level:%d ", rowstart, rowend, colstart, colend, level))
	//左 -》右
	for i := colstart; i <= colend; i++ {

		res[index] = matrix[rowstart][i]
		println(fmt.Sprintf("1row :%d col:%d index:%d", rowstart, i, index))
		index++

	}
	if index >= total {
		return
	}
	//上-》下
	for i := rowstart + 1; i <= rowend; i++ {
		// println(fmt.Sprintf("2row :%d col:%d index:%d ", i, colend, index))
		res[index] = matrix[i][colend]
		index++
	}
	if index >= total {
		return
	}
	//右-》左
	for i := colend - 1; i >= colstart; i-- {
		// println(fmt.Sprintf("3row :%d col:%d index:%d", rowend, i, index))
		res[index] = matrix[rowend][i]
		index++
	}
	if index >= total {
		return
	}
	//下-》上
	for i := rowend - 1; i >= rowstart+1; i-- {
		// println(fmt.Sprintf("4row :%d col:%d index:%d", i, colstart, index))
		res[index] = matrix[i][colstart]
		index++
	}
	level++
	loopmatrix(matrix, level, res, index)
}

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	col := len(matrix[0])
	println(fmt.Sprintf("row :%d col:%d ", row, col))
	res := make([]int, row*col)
	loopmatrix(matrix, 0, res, 0)
	return res
}

func main() {
	//input := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//input := [][]int{}
	input := [][]int{{1}}
	println(fmt.Sprintf("%#v", spiralOrder(input)))
}
