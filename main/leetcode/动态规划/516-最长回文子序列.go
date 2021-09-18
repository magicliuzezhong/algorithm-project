//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/18 09:41
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var word = "bbbab"
	var word = "abcabcabcabc"
	// "abcabcabcabc"
	var result = longestPalindromeSubseq(word)
	fmt.Println(result)
}

//
// longestPalindromeSubseq
// 这个动态规划问题毕竟经典，也比较难思考
//
func longestPalindromeSubseq(s string) int {
	// dp[i][j] = dp[i+1][j-1] && s[i] == s[j]  不跳过
	//          = dp[i+1][j-2] && s[i] == s[j]  跳过一个
	var wordLen = len(s)
	if wordLen == 0 {
		return 0
	}
	var dp = make([][]int, wordLen)
	for i := 0; i < wordLen; i++ {
		dp[i] = make([]int, wordLen)
		dp[i][i] = 1
	}
	for i := wordLen - 1; i >= 0; i-- {
		for j := i + 1; j < wordLen; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = maxFunc12(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][wordLen-1]
}

func maxFunc12(a, b int) int {
	if a > b {
		return a
	}
	return b
}
