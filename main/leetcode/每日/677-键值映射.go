package main

import "fmt"

// 一看到前缀第一时间就想到了前缀树
func main() {
	var mapSum = Constructor1()
	mapSum.Insert("apple", 3)
	mapSum.Insert("app", 2)
	var res1 = mapSum.Sum("appl") // return 3 (apple = 3)
	fmt.Println(res1)
	var res2 = mapSum.Sum("ap") // return 5 (apple + app = 3 + 2 = 5)
	fmt.Println(res2)
}

type Trie struct {
	word     string
	val      int
	isEnd    bool
	children [26]*Trie
}

type MapSum struct {
	trie *Trie
}

func Constructor1() MapSum {
	return MapSum{
		trie: &Trie{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	var trie = this.trie
	var keyLen = len(key)
	for i := 0; i < keyLen; i++ {
		var word = key[i : i+1]
		var wordSite = key[i] - 'a'
		var expectedVal = 0
		var isEnd = i == keyLen-1
		if isEnd {
			expectedVal = val
		}
		if trie.children[wordSite] == nil {
			trie.children[wordSite] = &Trie{
				word:  word,
				val:   expectedVal,
				isEnd: isEnd,
			}
		} else if isEnd {
			trie.children[wordSite].isEnd = true
			trie.children[wordSite].val = expectedVal
		}
		trie = trie.children[wordSite]
	}
}

func (this *MapSum) Sum(prefix string) int {
	var trie = this.trie
	var preLen = len(prefix)
	for i := 0; i < preLen; i++ {
		var wordSite = prefix[i] - 'a'
		if trie.children[wordSite] == nil {
			return 0
		}
		trie = trie.children[wordSite]
	}
	if trie == nil {
		return 0
	}
	var count = 0
	var stack = []*Trie{trie}
	for len(stack) > 0 {
		var top = stack[len(stack)-1]
		stack = stack[:len(stack)-1] //出栈
		if top.isEnd {
			count += top.val
		}
		var topChildren = top.children
		for _, child := range topChildren {
			if child != nil {
				stack = append(stack, child)
			}
		}
	}
	return count
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
