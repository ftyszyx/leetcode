package main

import (
	"fmt"
)

func main() {

	var dict = []string{"aaaa", "aaa"}
	var s = "aaaaaaa"
	fmt.Println(fmt.Sprintf("%t", wordBreak(s, dict)))
}

func wordBreak(s string, wordDict []string) bool {
	var dict = make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}
	var str = []byte(s)
	var n = len(str)
	if n == 0 {
		return true
	}
	//fmt.Println(fmt.Sprintf("n:%d ", n))
	var res = make([]bool, n+1)
	res[0] = true //res[i]代表 0到i没问题
	//求出
	for i := 0; i < n; i++ {
		for j := i; j < n && res[i]; j++ {
			substr := string(str[i : j+1])
			_, have := dict[substr]
			//fmt.Println(fmt.Sprintf("i:%d j:%d have:%t substr:%s", i, j, have, substr))
			if have == true {
				res[j+1] = have
			}

		}
	}
	return res[n]
}
