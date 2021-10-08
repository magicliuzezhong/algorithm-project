//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/27 16:04
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var s = "1*1"
	var result = numDecodings1(s)
	fmt.Println(result)
}

func numDecodings1(s string) int {
	// 1* 9
	// * 9
	// 2* 6
	// ** 15

	var sLen = len(s)
	var dp = make([]int, sLen)
	if s[0] == '*' {
		dp[0] = 9
	} else if s[0] != '0' {
		dp[0] = 1
	}
	var record = map[string]int{
		"1*": 9,
		"*":  9,
		"2*": 6,
		"**": 15,
	}
	const mod = 1e9 + 7
	for i := 1; i < sLen; i++ {
		if i == 1 {
			dp[i] = record[s[i-1:i]] % mod
		} else {
			var num = 0
			if s[i-2] == '*' && s[i-1] != '*' {
				if s[i-1] <= 6 {
					num = 2
				} else {
					num = 1
				}
			} else {
				num = record[s[i-2:i]]
			}
			var num2 = record[s[i-1:i]]
			dp[i] = (num*dp[i-2] + num2*dp[i-1]) % mod
		}
	}
	return dp[sLen-1]
}
