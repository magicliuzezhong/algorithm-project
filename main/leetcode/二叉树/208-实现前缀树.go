//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/19 17:56
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	trie := Constructor()
	var addWord = []string{"trie", "insert", "aaa", "aa", "search", "search", "startswith", "insert", "search"}
	for _, val := range addWord {
		trie.Insert(val)
	}
	var res = trie.StartsWith("word")
	fmt.Println(res)
	var result = trie.Search("insert")
	fmt.Println(result)
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	var node = this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) searchPrefix(prefix string) *Trie {
	var node = this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	var node = this.searchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}
