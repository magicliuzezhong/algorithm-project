//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/9 14:35
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	//var arr = []int{7,1,5,3,6,4}
	//var arr = []int{1,2,3,4,5}
	//var arr = []int{5,4,3,2,1}
	//var arr = []int{3, 2, 6, 5, 0, 3}
	var arr = []int{1, 9, 6, 9, 1, 7, 1, 1, 5, 9, 9, 9}
	var result = maxProfit(arr)
	fmt.Println(result)
}

//
// maxProfit
// 看了下题解，还是人才多啊，因为买卖不受限制，那么收集所有的上坡就行了
// 好的这就去银行开户(-_-)
//
func maxProfit(prices []int) int {
	var numLen = len(prices)
	var maxVal = 0
	for i := 1; i < numLen; i++ {
		if prices[i] > prices[i-1] {
			maxVal += prices[i] - prices[i-1]
		}
	}
	return maxVal
}

//
// maxProfit
// 没有看题解哦，一次性做出来了
//
func maxProfitBak(prices []int) int {
	// 7  1  5  3  6  4
	// prices[i - 1] < prices[i] -> 买入
	// prices[i - 1] > prices[i] -> 卖出
	var numLen = len(prices)
	var maxVal = 0
	var maxI = math.MinInt64
	var minI = prices[0]
	for i := 1; i < numLen; i++ {
		if prices[i] > prices[i-1] {
			maxI = prices[i]
			if i == numLen-1 && maxI > minI {
				maxVal += maxI - minI
			}
		} else if prices[i] < prices[i-1] {
			if maxI > minI {
				maxVal += maxI - minI
				maxI = math.MinInt64
			}
			minI = prices[i]
		} else if prices[i] == prices[i-1] && i == numLen-1 && maxI > minI {
			maxVal += maxI - minI
		}
	}
	return maxVal
}
