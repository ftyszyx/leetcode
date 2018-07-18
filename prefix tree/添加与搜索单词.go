package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.AddWord("at")
	obj.AddWord("and")
	obj.AddWord("an")
	obj.AddWord("bat")

	fmt.Println(fmt.Sprintf("%t", obj.Search(".at")))

}

type WordDictionary struct {
	childs [26]*WordDictionary
	isword bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	strbyte := []byte(word)
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
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {

	//fmt.Println(fmt.Sprintf("%s", word))
	//fmt.Println(fmt.Sprintf("%+v", this))
	if word == "" {
		return this.isword
	}
	strbyte := []byte(word)
	for index, char := range strbyte {
		if char == '.' {
			for _, child := range this.childs {
				if child != nil {
					if child.Search(string(strbyte[index+1:])) == true {
						return true
					}
				}
			}
			return false
		} else {
			child := this.childs[char-'a']
			if child != nil {
				return child.Search(string(strbyte[index+1:]))
			} else {
				//fmt.Println(fmt.Sprintf("not find %d", char))
				return false
			}
		}
	}
	return false
}
