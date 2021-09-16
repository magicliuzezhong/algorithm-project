//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/16 16:17
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	var result = uniquePathsWithObstacles(arr)
	fmt.Println(result)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var yLen = len(obstacleGrid)
	var xLen = len(obstacleGrid[0])
	var dp = make([][]int, yLen)
	for i := 0; i < yLen; i++ {
		dp[i] = make([]int, xLen)
	}
	for i := 0; i < xLen; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 0; i < yLen; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < yLen; i++ {
		for j := 1; j < xLen; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[yLen-1][xLen-1]
}
