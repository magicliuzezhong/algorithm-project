//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/13 16:24
// @Company cloud-ark.com
//
package main

import "fmt"

//   __       _______    ______
//  |  |     |__    /   |__   /
//  |  |        /  /      /  /
//  |  |       /  /      /  /
//  |  |___   /  /__    /  /__
//  |______| /______|  /______|

func main() {
	//var word = "applepenapple"
	//var wordDict = []string{"apple", "pen"}
	var word = "catsandog"
	var wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(word, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	// 转移方程
	// dp[i] = dp[j] && check(s[j...:i])
	// applepenapple

	var wordDictMap = make(map[string]bool)
	var maxWordLen = 0
	for _, word := range wordDict {
		wordDictMap[word] = true
		maxWordLen = maxFunc8(maxWordLen, len(word))
	}
	// 8 - 2 = 6  6 7 8
	var sLen = len(s)
	var dp = make([]bool, sLen+1)
	dp[0] = true
	var j = 0
	for i := 1; i <= sLen; i++ {
		if i-maxWordLen < 0 {
			j = 0
		} else {
			j = i - maxWordLen
		}
		for ; j >= 0 && j < i; j++ { //最大步进最大单词的word
			if dp[j] && wordDictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[sLen]
}

func maxFunc8(a, b int) int {
	if a > b {
		return a
	}
	return b
}
