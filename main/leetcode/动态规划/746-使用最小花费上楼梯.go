//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/3 09:14
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{10, 15, 20, 12, 13, 11}
	var result = minCostClimbingStairs(arr)
	fmt.Println(result)
}

func minCostClimbingStairs(cost []int) int {
	var arrLen = len(cost)
	var dp = make([]int, arrLen+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= arrLen; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[arrLen]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
