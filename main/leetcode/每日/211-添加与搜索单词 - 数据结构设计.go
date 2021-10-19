//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/19 21:08:31
// @Company cloud-ark.com
//
package main

import "fmt"

// 前缀树，用于快速搜索
func main() {
	var words = []string{"worddictionary", "bacd", "addword", "addword", "search", "search", "search", "search"}
	var searchWords = []string{"", "bad", "dad", "mad", "pad", "bad", ".ad", "b.."}
	var w = Constructor()
	for _, val := range words {
		w.AddWord(val)
	}
	for _, val := range searchWords {
		var result = w.Search(val)
		fmt.Println(result)
	}
}

type trieNode struct {
	children [26]*trieNode
	isEnd    bool
}

func (t *trieNode) addNode(word string) {
	var node = t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &trieNode{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

type WordDictionary struct {
	trieNode *trieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		trieNode: &trieNode{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.trieNode.addNode(word)
}

func (this *WordDictionary) Search(word string) bool {
	var node = this.trieNode
	var wordLen = len(word)
	var searchFunc func(index int, node *trieNode) bool
	searchFunc = func(index int, node *trieNode) bool {
		if index == wordLen {
			return node.isEnd
		}
		if word[index] == '.' {
			for _, trie := range node.children {
				if trie != nil && searchFunc(index+1, trie) {
					return true
				}
			}
		} else {
			var nextNode = node.children[word[index]-'a']
			if nextNode != nil && searchFunc(index+1, nextNode) {
				return true
			}
		}
		return false
	}
	return searchFunc(0, node)
}
