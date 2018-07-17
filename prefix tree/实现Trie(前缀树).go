package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Insert("test")
	param_2 := obj.Search("test")
	param_3 := obj.StartsWith("aes")

	fmt.Printf(" param_2:%t param_3:%t", param_2, param_3)
	//fmt.Printf("obj:%+v param_2:%b param_3:%b", obj, param_2, param_3)
}

type Trie struct {
	childs [26]*Trie
	isword bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	strbyte := []byte(word)
	curroot := this
	for _, char := range strbyte {
		curchild := curroot.childs[char-'a']
		if curchild == nil {
			temp := Constructor()
			curchild = &temp

			//fmt.Printf("set:%+v char:%d", curchild, char)
		}
		curroot.childs[char-'a'] = curchild
		curroot = curchild
	}
	curroot.isword = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	strbyte := []byte(word)
	curroot := this
	for _, char := range strbyte {
		curroot = curroot.childs[char-'a']
		if curroot == nil {
			return false
		}
	}
	return curroot.isword
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	strbyte := []byte(prefix)
	curroot := this
	for _, char := range strbyte {
		curroot = curroot.childs[char-'a']
		if curroot == nil {
			return false
		}
	}
	return true
}
