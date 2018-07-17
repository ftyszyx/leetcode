package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Insert("apple", 3)
	num1 := obj.Sum("app")
	obj.Insert("app", 2)
	num2 := obj.Sum("ap")
	fmt.Println(fmt.Sprintf(" num1:%d num2:%d", num1, num2))
	fmt.Println(fmt.Sprintf("obj:%+v", obj))
}

type MapSum struct {
	childs [26]*MapSum
	isword bool
	num    int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	strbyte := []byte(key)
	curroot := this
	for _, char := range strbyte {
		curchild := curroot.childs[char-'a']
		if curchild == nil {
			temp := Constructor()
			curchild = &temp
		}
		curroot.childs[char-'a'] = curchild
		curroot = curchild
	}
	curroot.isword = true
	curroot.num = val
}

func getchildNum(root *MapSum) int {
	if root == nil {
		return 0
	}
	var sum = 0
	if root.isword == true {
		//fmt.Println(fmt.Sprintf("get world root:%+v", root))
		sum += root.num
	}
	for _, value := range root.childs {
		sum += getchildNum(value)
	}

	return sum
}

func (this *MapSum) Sum(prefix string) int {
	strbyte := []byte(prefix)
	curroot := this
	for _, char := range strbyte {
		curroot = curroot.childs[char-'a']
		if curroot == nil {
			return 0
		}
	}
	var sum = 0
	sum += getchildNum(curroot)
	return sum
}
