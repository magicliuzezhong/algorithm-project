//
// Package main
// @Description：
// @Author：liuzezhong 2021/11/12
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var n = 1000
	var k = 1000
	var res = kInversePairs(n, k)
	fmt.Println(res)
	var res1 = kInversePairs1(n, k)
	fmt.Println(res1)
}

func kInversePairs(n int, k int) int {
	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[1][0] = 1
	var mod = 1000000007
	for i := 2; i <= n; i++ {
		for j := 0; j <= i*(i-1)/2 && j <= k; j++ {
			if j == 0 {
				dp[i][j] = 1
			} else {
				var start = 0
				if j-i+1 > start {
					start = j - i + 1
				}
				for x := start; x <= j; x++ {
					dp[i][j] = (dp[i][j] + dp[i-1][x]) % mod
				}
			}
		}
	}
	return dp[n][k]
}

func kInversePairs1(n int, k int) int {
	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[1][0] = 1
	var mod = 1000000007
	for i := 2; i <= n; i++ {
		for j := 0; j <= i*(i-1)/2 && j <= k; j++ {
			// dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-i]
			dp[i][j] = dp[i-1][j]
			if j >= 1 {
				dp[i][j] += dp[i][j-1]
			}
			if j >= i {
				dp[i][j] -= dp[i-1][j-i]
			}
			dp[i][j] = (dp[i][j] + mod) % mod
		}
	}
	return dp[n][k]
}
