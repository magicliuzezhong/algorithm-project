//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/3 09:04
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var result = climbStairs(4)
	// 1 1 1 1
	// 2 1 1
	// 1 2 1
	// 1 1 2
	// 2 2
	fmt.Println(result)
}

//
// climbStairs
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
func climbStairs(n int) int {
	// 动态规划，每次走1、2个台阶，
	// dp[n] = dp[n - 1] + dp[n - 2]
	if n < 3 {
		return n
	}
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
