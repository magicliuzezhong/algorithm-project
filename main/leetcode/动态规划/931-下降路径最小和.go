//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/14 15:07
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	var arr = [][]int{
		{2, 1, 3},
		{6, 5, 4},
		//{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(arr))
}

//
// minFallingPathSum
// 原地修改就行，当然有的题解说这是一种流氓玩法，自己怎么舒服怎么来？
//
func minFallingPathSum(matrix [][]int) int {
	var yLen = len(matrix)
	var xLen = len(matrix[0])
	var val = math.MaxInt64
	var tmpMin = 0
	for i := 0; i < yLen; i++ {
		for j := 0; j < xLen; j++ {
			tmpMin = 0
			if i > 0 {
				tmpMin = matrix[i-1][j]
				if j > 0 {
					tmpMin = minFunc9(matrix[i-1][j-1], tmpMin)
				}
				if j+1 < xLen {
					tmpMin = minFunc9(matrix[i-1][j+1], tmpMin)
				}
			}
			matrix[i][j] = matrix[i][j] + tmpMin
			if i == yLen-1 {
				val = minFunc9(val, matrix[i][j])
			}
		}
	}
	return val
}

func minFunc9(a, b int) int {
	if a > b {
		return b
	}
	return a
}
