//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/14 15:29
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	var arr = [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
		{4, 10, 8, 3, 1},
	}
	fmt.Println(minimumTotal(arr))
}

//
// minimumTotal
// 跟 931-下降路径最小和 一模一样的算法
//
func minimumTotal(triangle [][]int) int {
	var yLen = len(triangle)
	var lastXlen = 0
	var val = math.MaxInt64
	for i := 0; i < yLen; i++ {
		var xLen = len(triangle[i])
		lastXlen = xLen - 1
		for j := 0; j < xLen; j++ {
			if lastXlen > 0 {
				if j == 0 {
					triangle[i][j] = triangle[i][j] + triangle[i-1][0]
				}
				if j == lastXlen {
					triangle[i][j] = triangle[i][j] + triangle[i-1][lastXlen-1]
				}
				if j > 0 && j < lastXlen {
					triangle[i][j] = triangle[i][j] + minFunc10(triangle[i-1][j-1], triangle[i-1][j])
				}
			}
			if i == yLen-1 {
				val = minFunc10(val, triangle[i][j])
			}
		}
	}
	return val
}

func minFunc10(a, b int) int {
	if a > b {
		return b
	}
	return a
}
