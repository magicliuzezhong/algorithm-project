//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/16 15:15
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var m = 32
	var n = 12
	var result = uniquePaths(m, n)
	fmt.Println(result)
	//var result2 = uniquePaths2(m, n)
	//fmt.Println(result2)
}

func uniquePaths(m int, n int) int {
	// dp[i][j] = dp[i - 1][j] * dp[i][j - 1]
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePaths2(m int, n int) int {
	var result = 0
	backtrace(m, n, 1, 1, &result)
	return result
}

func backtrace(m, n int, currentY, currentX int, result *int) {
	if m == currentY && n == currentX {
		*result = *result + 1
		return
	}
	if currentY+1 <= m {
		backtrace(m, n, currentY+1, currentX, result)
	}
	if currentX+1 <= n {
		backtrace(m, n, currentY, currentX+1, result)
	}
}
