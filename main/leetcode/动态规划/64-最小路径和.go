//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/2 11:22
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

//
// 给定一个包含非负整数的 mxn网格grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//   0 1 2
// 0 1 3 1
// 1 1 5 1
// 2 4 2 1
// 说明：每次只能向下或者向右移动一步。
//
func main() {
	//var grid = [][]int{
	//	{1, 3, 1},
	//	{1, 5, 1},
	//	{4, 2, 1},
	//}
	var grid = [][]int{
		{1, 4, 8, 6, 2, 2, 1, 7},
		{4, 7, 3, 1, 4, 5, 5, 1},
		{8, 8, 2, 1, 1, 8, 0, 1},
		{8, 9, 2, 9, 8, 0, 8, 9},
		{5, 7, 5, 7, 1, 8, 5, 5},
		{7, 0, 9, 4, 5, 6, 5, 6},
		{4, 9, 9, 7, 9, 1, 9, 0},
	}

	var result = minPathSum(grid)
	fmt.Println(result)
	//var result1 = minPathSum1(grid)
	//fmt.Println(result1)
}

//
// minPathSum
// 虽然说之前做过一次了，再做一次也无妨
//
func minPathSum(grid [][]int) int {
	var yLen = len(grid)
	var xLen = len(grid[0])
	for i := 1; i < yLen; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < xLen; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < yLen; i++ {
		for j := 1; j < xLen; j++ {
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] += grid[i][j-1]
			} else {
				grid[i][j] += grid[i-1][j]
			}
		}
	}
	return grid[yLen-1][xLen-1]
}

func minPathSum1(grid [][]int) int {
	var xLen = len(grid)
	var yLen = len(grid[0])
	var dp = make([][]int, xLen)
	for i := 0; i < xLen; i++ {
		dp[i] = make([]int, yLen)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < xLen; i++ {
		for j := 0; j < yLen; j++ {
			var min = math.MaxInt64
			if i-1 >= 0 {
				min = dp[i-1][j]
			}
			if j-1 >= 0 {
				min = int(math.Min(float64(min), float64(dp[i][j-1])))
			}
			if min != math.MaxInt64 {
				dp[i][j] = grid[i][j] + min
			}
		}
	}
	return dp[xLen-1][yLen-1]
}
