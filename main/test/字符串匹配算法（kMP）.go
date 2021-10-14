//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/14 10:00
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	var word = "aaaabaaabaaaaag"
	var target = "aaaaa"
	//var res = IndexOf(word, target)
	var res = kmp(word, target)
	fmt.Println(res)
}

//
// kmp
// 和朴素字符串匹配相比，其跳过的节点更多，
// 因为根据ababc可知，如果出现某个字符已经不匹配的情况下，我们能直接跳到之前匹配后相匹配的位置
// 从而避免了朴素字符串那样每一个都要对比一遍的情况，减少了匹配次数。
//
func kmp(word, target string) int {
	var wordLen = len(word)
	var targetLen = len(target)
	switch {
	case targetLen == 0:
		return 0
	case targetLen > wordLen:
		return -1
	case targetLen == wordLen:
		if word == target {
			return 0
		}
		return -1
	default:
		var nextArr = make([]int, targetLen)
		for i := 0; i < targetLen; i++ {
			var t = 0
			for j := 0; j < i-1; j++ {
				if target[:j+1] == target[i-(j+1):i] {
					t = j + 1
				}
			}
			nextArr[i] = t
		}
		var wordIndex, targetIndex = 0, 0
		for wordIndex < wordLen && targetIndex < targetLen {
			for targetIndex < targetLen && word[wordIndex] == target[targetIndex] {
				wordIndex++
				targetIndex++
			}
			if targetIndex == targetLen {
				return wordIndex - targetLen
			}
			wordIndex = wordIndex - nextArr[targetIndex] + 1
			targetIndex = 0
		}
		fmt.Println(nextArr)
		return -1
	}
}

//
// IndexOf
// 朴素字符串匹配算法
//
func IndexOf(word, target string) int {
	var wordLen = len(word)
	var targetLen = len(target)
	switch {
	case targetLen == 0:
		return 0
	case targetLen > len(word):
		return -1
	case targetLen == len(word):
		if word == target {
			return 0
		}
		return -1
	default:
		var wordIndex = 0
		var targetIndex = 0
		for wordIndex < wordLen {
			if wordLen-wordIndex < targetLen {
				return -1
			}
			for targetIndex < targetLen && word[wordIndex] == target[targetIndex] {
				wordIndex++
				targetIndex++
			}
			if targetIndex == targetLen {
				return wordIndex - targetLen
			}
			wordIndex = wordIndex - targetIndex + 1
			targetIndex = 0
		}
		return -1
	}
}
