//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/23 15:42
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	var coins = []int{11}
	var target = 11
	var result = coinChange(coins, target)
	fmt.Println(result)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	var dp = make([]int, amount+1)
	for _, coin := range coins {
		if coin <= amount {
			dp[coin] = 1
		}
	}
	var tmp = 0
	for i := 1; i <= amount; i++ {
		tmp = math.MaxInt64
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				tmp = minFunc16(dp[i-coin]+1, tmp)
			}
		}
		if tmp == math.MaxInt64 {
			tmp = -1
		}
		dp[i] = tmp
	}
	return dp[amount]
}

func minFunc16(a, b int) int {
	if a > b {
		return b
	}
	return a
}
