//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/14 09:49
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var s = "22616"
	fmt.Println(numDecodings(s))
}

// 1  A
// 2  B
// ...
// 26 Z
func numDecodings(s string) int {
	// 使用一种解码方式   dp[i] = dp[i -1]
	// 使用两种解码方式   dp[i] = dp[i - 2]
	var sLen = len(s)
	var dp = make([]int, sLen+1)
	dp[0] = 1
	for i := 1; i <= sLen; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			dp[i] += dp[i-2]
		}
	}
	return dp[sLen]
}
