//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/17 14:18
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	var arr = [][]byte{
		//{'0'},
		//{'1', '1'},
		//{'1', '0'},
		{'1', '0', '1', '0', '0'},
		{'1', '0', '0', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '1'},

		//{'1', '1', '2', '2', '2'},
		//{'2', '2', '4', '5', '6'},
		//{'3', '4', '7', '9', '11'},
		//{'4', '5', '8', '11', '13'},
	}
	// 13 - 3 - 2 + 1 = 8
	fmt.Println(maximalSquare(arr))
	fmt.Println(maximalSquare1(arr))
}

//
// maximalSquare
// 虽然已经做出来了，但是不妨把官方的思路也跟着实现下
//
func maximalSquare(matrix [][]byte) int {
	var yLen = len(matrix)
	var xLen = len(matrix[0])
	var dp = make([][]int, yLen)
	var maxSide = 0
	for i := 0; i < yLen; i++ {
		dp[i] = make([]int, xLen)
		for j := 0; j < xLen; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				maxSide = 1
			}
		}
	}
	for i := 1; i < yLen; i++ {
		for j := 1; j < xLen; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = minFunc12(dp[i-1][j], minFunc12(dp[i][j-1], dp[i-1][j-1])) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}

//
// maximalSquare
// 我自己想到了解题的办法
// 采用的是二维数组前缀和
// 官方采用的动态规划
// 流下了不争气的泪水(-_-)
// 当前空间复杂度为O(m*n)
// 我又想了想能否优化下空间复杂度呢，其实是可以的，完全可以再原地置换修改matrix数组，
// 可以个锤子，byte置换不了
//
func maximalSquare1(matrix [][]byte) int {
	var yLen = len(matrix)
	var xLen = len(matrix[0])
	var maxVal = 0
	var sum = make([][]int, yLen)
	for i := 0; i < yLen; i++ {
		sum[i] = make([]int, xLen)
	}
	sum[0][0] = change2Num(matrix[0][0])
	if sum[0][0] == 1 {
		maxVal = 1
	}
	for i := 1; i < yLen; i++ {
		sum[i][0] += sum[i-1][0] + change2Num(matrix[i][0])
		if maxVal == 0 && sum[i][0] > 0 {
			maxVal = 1
		}
	}
	for i := 1; i < xLen; i++ {
		sum[0][i] += sum[0][i-1] + change2Num(matrix[0][i])
		if maxVal == 0 && sum[0][i] > 0 {
			maxVal = 1
		}
	}
	var maxBorder = 2
	var left = 0
	var top = 0
	var val = 0
	for i := 1; i < yLen; i++ {
		for j := 1; j < xLen; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + change2Num(matrix[i][j])
			if maxVal == 0 && sum[i][j] > 0 {
				maxVal = 1
			}
			top = i - maxBorder
			left = j - maxBorder
			if top >= -1 && left >= -1 {
				val = sum[i][j]
				if left >= 0 {
					val -= sum[i][left]
				}
				if top >= 0 {
					val -= sum[top][j]
				}
				if left >= 0 && top >= 0 {
					val += sum[top][left]
				}
				if val == maxBorder*maxBorder {
					maxVal = val
					maxBorder++
				}
			}
		}
	}
	return maxVal
}

func change2Num(b byte) int {
	if b == '0' {
		return 0
	}
	return 1
}

func minFunc12(a, b int) int {
	if a > b {
		return b
	}
	return a
}
