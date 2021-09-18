//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/18 09:03
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var word = "babab"
	var word = "aaaaaa"

	// a a a a a
	// 1 1 1 1 1
	// 1 2 2 4

	// a a
	// a b
	// a a a

	fmt.Println(longestPalindrome(word))
}

//
// longestPalindrome
// 比较通常的做法
//
func longestPalindrome(s string) string {
	var wordLen = len(s)
	if wordLen == 0 {
		return ""
	}
	var dp = make([][]bool, wordLen)
	for i := 0; i < wordLen; i++ {
		dp[i] = make([]bool, wordLen)
		dp[i][i] = true
	}
	// dp[i][j] = dp[i + 1][j - 1] && s[i] == s[j]
	var maxVal = 0
	var maxWord = s[0:1]
	for i := wordLen - 1; i >= 0; i-- {
		for j := i + 1; j < wordLen; j++ {
			if j == i+1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] && j-i+1 > maxVal {
				maxVal = j - i + 1
				maxWord = s[i : j+1]
			}
		}
	}
	return maxWord
}
