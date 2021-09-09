//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/9 14:06
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

//
// 给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
//
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//
func main() {
	var arr = []int{7, 1, 5, 3, 6, 4}
	var result = maxProfit(arr)
	fmt.Println(result)
}

//
// maxProfit
// 其实和1014题一模一样的，前面记录最低点，用当前值减去最低点就得到了结果
// 这道题我没有看题解哦 oh yes
//
func maxProfit1(prices []int) int {
	var numLen = len(prices)
	var maxVal = 0
	var minI = prices[0]
	for i := 1; i < numLen; i++ {
		maxVal = maxFunc5(maxVal, prices[i]-minI)
		minI = minFunc5(minI, prices[i])
	}
	return maxVal
}

func minFunc5(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxFunc5(a, b int) int {
	if a > b {
		return a
	}
	return b
}
