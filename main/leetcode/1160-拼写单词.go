//
// Package main
// @Description：
// @Author：liuzezhong 2021/7/29 17:47
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var words = []string{"cat", "bt", "hat", "tree"}
	var chars = "atach"
	//var words = []string{"hello", "world", "leetcode"}
	//var chars = "welldonehoneyr"
	//var chars = ""
	fmt.Println(countCharacters(words, chars))
}

func countCharacters(words []string, chars string) int {
	var wordMap = createMap(chars)
	var result = 0
	for _, val := range words {
		var valMap = createMap(val)
		var count = 0
		for key, value := range  valMap {
			var charCount, ok = wordMap[key]
			if ok && value <= charCount {
				count += value
			}
		}
		if count == len(val) {
			result += count
		}
	}
	return result
}

func createMap(chars string) map[string]int {
	var wordMap = make(map[string]int)
	var charLen = len(chars)
	for i := 0; i < charLen; i++ {
		var word = chars[i : i+1]
		var count, ok = wordMap[word]
		if ok {
			wordMap[word] = count + 1
		} else {
			wordMap[word] = 1
		}
	}
	return wordMap
}
