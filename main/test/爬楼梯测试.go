//
// Package test
// @Description：
// @Author：liuzezhong 2021/8/5 17:29
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {

	// 分硬币问题  一共有 1 3 5种硬币，给定金额为10，请问最小分法是什么
	fmt.Println(gets(19))
}

func gets(target int) int {
	var dp = make([]int, target+1)

	// dp[n] = math.Min(dp[n - 1], dp[n - 3], dp[n -5]) + 1
	for i := 1; i <= target; i++ {
		if i == 1 || i == 3 || i == 5 {
			dp[i] = 1
			continue
		}
		if i-1 >= 1 {
			dp[i] = dp[i-1]
		}
		if i-3 >= 1 {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-3])))
		}
		if i-5 >= 1 {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-5])))
		}
		dp[i] += 1
	}
	return dp[target]
}
