package main

import (
	"fmt"
	"strings"
)

func main() {
	var dict = []string{"cat", "bat", "rat"}
	var sentence = "the cattle was rattled by the battery"
	fmt.Println(fmt.Sprintf("%s", replaceWords(dict, sentence)))
}

func replaceWords(dict []string, sentence string) string {
	obj := Constructor()
	for _, str := range dict {
		obj.Insert(str)
	}
	// var buffer bytes.Buffer
	arrstr := strings.Split(sentence, " ")
	for index, item := range arrstr {
		itemstr := strings.TrimSpace(item)
		arrstr[index] = obj.Search(itemstr)

	}
	return strings.Join(arrstr, " ")
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
func (this *Trie) Search(word string) string {
	strbyte := []byte(word)
	curroot := this

	for index, char := range strbyte {
		curroot = curroot.childs[char-'a']
		if curroot == nil {
			return word
		} else {
			if curroot.isword {
				return string(strbyte[:index+1])
			}
		}
	}
	return word
}
