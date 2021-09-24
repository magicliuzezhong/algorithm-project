//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/23 15:57
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var coins = []int{1, 2, 5}
	//var target = 5
	var coins = []int{2}
	var target = 1
	var result = change(target, coins)
	fmt.Println(result)
}

//
// change
// 学到一个词NP难问题
// 什么意思呢，
// 一个无法在多项式时间复杂度内就行求解的
// 并且可以再多项式时间复杂度验证的问题
// 只能通过 边遍历 边验证 进行求解的问题
//
func change(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
