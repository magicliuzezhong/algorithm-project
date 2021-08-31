//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/20 17:13
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var result = longestValidParentheses(")(((((()())()()))()(()))(")
	//var result = longestValidParentheses("()(())")
	//var result = longestValidParentheses("()")
	fmt.Println(result)
}

//
// longestValidParentheses
// 如果以 "(" 结尾那么 dp[i] = 0
// 如果以 ")" 结尾， 那么判断是否大于2
// 不大于 2 dp[i] = 2
// 不于 2 dp[i] = dp[i - dp[i - 2]] + 2
//
func longestValidParentheses(s string) int {
	var sLen = len(s)
	var maxLength = 0
	var dp = make([]int, sLen)
	for i := 1; i < sLen; i++ {
		if s[i:i+1] == ")" {
			if s[i-1:i] == "(" { //如果"("
				if i == 1 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
			} else {
				if dp[i-1] > 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1:i-dp[i-1]] == "(" {
					dp[i] = dp[i-1] + 2
					if i-dp[i-1]-2 >= 0 {
						dp[i] += dp[i-dp[i-1]-2]
					}
				}
			}
			if dp[i] > 0 && dp[i] > maxLength {
				maxLength = dp[i]
			}
		}
	}
	return maxLength
}
