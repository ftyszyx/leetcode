package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type NumSlice []int

func (this NumSlice) Len() int { return len(this) }
func (this NumSlice) Less(i, j int) bool {
	astr := strconv.Itoa(this[i])
	bstr := strconv.Itoa(this[j])
	ab, _ := strconv.Atoi(astr + bstr)
	ba, _ := strconv.Atoi(bstr + astr)
	return ab-ba > 0
}

func (this NumSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func largestNumber(nums []int) string {
	sort.Sort(NumSlice(nums))
	//println(fmt.Sprintf("num:%#v", nums))
	var res bytes.Buffer
	for _, item := range nums {
		res.WriteString(strconv.Itoa(item))
	}
	resstr := strings.TrimLeft(res.String(), "0")
	if resstr == "" {
		resstr = "0"
	}
	return resstr
}

func main() {
	input := []int{3, 30, 34, 5, 9}
	println(fmt.Sprintf("%s", largestNumber(input)))
}
