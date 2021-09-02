//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/2 10:55
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var n = 25
	var result = tribonacci(n)
	var result1 = tribonacci1(n)
	fmt.Println(result)
	fmt.Println(result1)
}

//
// tribonacci
// 泰波那契序列Tn定义如下：
//
// T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0的条件下 Tn+3 = Tn + Tn+1 + Tn+2
//
// 给你整数n，请返回第 n 个泰波那契数Tn 的值。
//
func tribonacci(n int) int {
	var numLen = n + 1
	if n < 3 {
		numLen = 3
	}
	var dp = make([]int, numLen)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func tribonacci1(n int) int {
	var modNum = n % 3
	var a = 0 // 1
	var b = 1 // 2
	var c = 1 // 3
	for i := 3; i <= n; i++ {
		modNum = i % 3
		if modNum == 1 {
			b = a + b + c
		} else if modNum == 2 {
			c = a + b + c
		} else {
			a = a + b + c
		}
	}
	if modNum == 0 {
		return a
	} else if modNum == 1 {
		return b
	} else {
		return c
	}
}
