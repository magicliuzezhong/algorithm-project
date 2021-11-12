//
// Package main
// @Description：
// @Author：liuzezhong 2021/11/12
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	var n = 1
	var res = getMoneyAmount(n)
	fmt.Println(res)
}

//
// 写一下解题思路
// 写啥啊，一开始就错了，一开始就想用二分法就行求解，
// 看了题解才恍然大悟，该题意在使用最少的金币，二分法意在最小的猜测次数
// 所以归根结底还是对二分法的最小猜测次数不熟练导致该题误入绝境（此处需要进行反思至少1秒钟）
// 话说回来，该题使用动态规划
// dp[i][j]表示在i和j区间花费最少金币的数量
// 假设 k 在i到j之间 k <i, j>
// 如果k没有猜中，那么可以得到规划方程
// dp[i][j] = k + max(dp[i][k-1], dp[k+1][j])
// 所以求出dp[0][n]即可
//
func getMoneyAmount(n int) int {
	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			var minCost = math.MaxInt64
			for k := i; k < j; k++ {
				var cost = k + max(dp[i][k-1], dp[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			dp[i][j] = minCost
		}
	}
	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
