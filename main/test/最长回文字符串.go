//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/9 14:53
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var words = "adaadashikghfiugisggaghdghshjaghhjd"
	var l = maxRevertStr(words)
	fmt.Println(l)
}

func maxRevertStr(word string) int {
	var wordLen = len(word)
	var dp [][]bool
	for i := 0; i < wordLen; i++ {
		var arr = make([]bool, wordLen)
		dp = append(dp, arr)
	}
	for i := 0; i < wordLen; i++ {
		dp[i][i] = true
	}
	var max = 0
	var maxWord = ""
	for i := wordLen - 1; i >= 0; i-- {
		for j := i + 1; j < wordLen; j++ {
			if i+1 == j {
				dp[i][j] = word[i:i+1] == word[j:j+1]
			} else {
				dp[i][j] = dp[i+1][j-1] && word[i:i+1] == word[j:j+1]
			}
			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
				maxWord = word[i : j+1]
			}
		}
	}
	fmt.Println(maxWord)
	return max
}
