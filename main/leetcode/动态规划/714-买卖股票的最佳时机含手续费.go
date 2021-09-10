//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/10 18:05
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var arr = []int{1, 3, 2, 8, 4, 9}
	var arr = []int{1, 3, 7, 5, 10, 3}
	//var fee = 2
	var fee = 3
	fmt.Println(maxProfit3(arr, fee))
}

//
// maxProfit3
// 哈哈，没看题解(-_-)
// 我又会了...
//
func maxProfit3(prices []int, fee int) int {
	// dp[i][0] 持有
	// dp[i][1] 不持有

	// dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] - prices[i])
	// dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] + prices[i] - free)

	var numLen = len(prices)
	var dp = make([][]int, numLen)
	for i := 0; i < numLen; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < numLen; i++ {
		dp[i][0] = maxFunc7(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = maxFunc7(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return dp[numLen-1][1]
}

func maxFunc7(a, b int) int {
	if a > b {
		return a
	}
	return b
}
