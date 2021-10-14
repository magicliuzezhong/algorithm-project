//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/15 17:25
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(7))
}

func nthUglyNumber(n int) int {
	var dp = make([]int, n+1)
	dp[1] = 1
	var p2, p3, p5 = 1, 1, 1
	var num1, num2, num3 = 0, 0, 0
	for i := 2; i <= n; i++ {
		num1 = dp[p2] * 2
		num2 = dp[p3] * 3
		num3 = dp[p5] * 5
		dp[i] = minFunc11(num1, minFunc11(num2, num3))
		if dp[i] == num1 {
			p2++
		}
		if dp[i] == num2 {
			p3++
		}
		if dp[i] == num3 {
			p5++
		}
	}
	return dp[n]
}

//
// nthUglyNumber
// 说实话，愣是没看懂题目说的啥意思
//
func nthUglyNumber1(n int) int {
	var dp = make([]int, n+1)
	dp[1] = 1
	var p2, p3, p5 = 1, 1, 1
	var num1, num2, num3 = 0, 0, 0
	for i := 2; i <= n; i++ {
		// 1 2 3 4 5 6 8 9 10
		num1 = dp[p2] * 2
		num2 = dp[p3] * 3
		num3 = dp[p5] * 5
		dp[i] = minFunc11(num1, minFunc11(num2, num3))
		if dp[i] == num1 {
			p2++
		}
		if dp[i] == num2 {
			p3++
		}
		if dp[i] == num3 {
			p5++
		}
	}
	return dp[n]
}

func minFunc11(a, b int) int {
	if a > b {
		return b
	}
	return a
}
