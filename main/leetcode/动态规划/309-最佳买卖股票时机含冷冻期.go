//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/10 17:08
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit2(arr))
}

//
// maxProfit2
// 看了下题解，非常经典的动态规划问题，所以我把122题的又以动态规划方式写了一遍，
// 总结：我脑子瓦特了(-_-) 不是在看题解的路上就是处于去看题解路上
//
func maxProfit2(prices []int) int {
	// dp[i][0] 当前持有股票
	// dp[i][1] 当前处于冷冻期
	// dp[i][2] 当前未持有股票并且不处于冷冻期

	// 当前持有股票 = max(上一天也持有, 上一天没有持有股票且不处于冷冻期 - 今天买入的负收益)
	// dp[i][0] = max(dp[i - 1][0], dp[i - 1][2] - price[i])
	// 当前处于冷冻期 = 上一天持有股票 + 卖出的收益
	// dp[i][1] = dp[i - 1][0] + price[i]
	// 当前未持有股票并且不处于冷冻期 = max(上一天处于冷冻期, 上一天也不持有股票并且不处于冷冻期)
	// dp[i][2] = max(dp[i - 1][1], dp[i - 1][2])
	var numLen = len(prices)
	var dp = make([][]int, numLen)
	for i := 0; i < numLen; i++ {
		dp[i] = make([]int, 3)
	}
	// 初始化
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < numLen; i++ {
		dp[i][0] = maxFunc6(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = maxFunc6(dp[i-1][1], dp[i-1][2])
	}
	return maxFunc6(dp[numLen-1][1], dp[numLen-1][2])
}

func maxFunc6(a, b int) int {
	if a > b {
		return a
	}
	return b
}
