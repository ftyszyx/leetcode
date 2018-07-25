package main

import (
	"fmt"
)

func main() {

	var dict = []string{"cat", "cats", "and", "sand", "dog"}
	var s = "catsanddog"
	fmt.Println(fmt.Sprintf("%#v", wordBreak(s, dict)))
}

type splitData struct {
	world     string
	lastindex int
}

type pointdata struct {
	info   []splitData
	isword bool
}

var worlds []string

func wordBreak(s string, wordDict []string) []string {
	worlds = []string{}
	var dict = make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}
	var str = []byte(s)
	var n = len(str)
	if n == 0 {
		return worlds
	}
	//fmt.Println(fmt.Sprintf("n:%d ", n))
	var res = make([]pointdata, n+1)
	res[0].isword = true //res[i]代表 0到i没问题
	//求出
	for i := 0; i < n; i++ {
		for j := i; j < n && res[i].isword; j++ {
			substr := string(str[i : j+1])
			_, have := dict[substr]

			//fmt.Println(fmt.Sprintf("i:%d j:%d have:%t substr:%s", i, j, have, substr))
			if have == true {

				res[j+1].info = append(res[j+1].info, splitData{substr, i})
				res[j+1].isword = have
			}

		}
	}
	if res[n].isword == false {
		return worlds
	} else {
		for _, item := range res[n].info {
			//fmt.Println(fmt.Sprintf("after %s add item:%s", item.world, item.world))
			initword(item.world, item.lastindex, res)
		}
		return worlds
	}

}

func initword(res string, curindex int, dataarr []pointdata) {
	if curindex == 0 {
		//fmt.Println(fmt.Sprintf("get string:%s\n", res))
		worlds = append(worlds, res)
	} else {
		for _, item := range dataarr[curindex].info {
			getstr := item.world + " " + res
			//fmt.Println(fmt.Sprintf("after %s add item:%s", getstr, item.world))
			initword(getstr, item.lastindex, dataarr)
		}
	}
}
