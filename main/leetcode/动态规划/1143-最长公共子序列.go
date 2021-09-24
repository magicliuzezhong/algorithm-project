//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/23 11:01
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var text1 = "aabcde"
	//var text2 = "acefae"
	var text1 = "jgtargjctqvijshexyjcjcre"
	var text2 = "pyzazexujqtsjebcnadahobwf"
	var result = longestCommonSubsequence(text1, text2)
	fmt.Println(result)
}

//
// longestCommonSubsequence
// 非常经典的动态规划问题
//
func longestCommonSubsequence(text1 string, text2 string) int {
	var text2Len = len(text2)
	if text2Len == 0 {
		return 0
	}
	var text1Len = len(text1)
	var dp = make([][]int, text1Len+1)
	for i := 0; i <= text1Len; i++ {
		dp[i] = make([]int, text2Len+1)
	}
	// dp[i][j] = dp[i-1][j-1] + 1   text1[i] == text2[j]
	// dp[i][j] = max(dp[i-1][j], dp[i][j-1])  text1[i] != text2[j]
	for i := 0; i < text1Len; i++ {
		for j := 0; j < text2Len; j++ {
			if text1[i:i+1] == text2[j:j+1] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = maxFunc15(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[text1Len][text2Len]
}

func maxFunc15(a, b int) int {
	if a > b {
		return a
	}
	return b
}
