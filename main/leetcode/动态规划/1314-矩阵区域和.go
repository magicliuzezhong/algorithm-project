//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/15 15:39
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	// {1, 2, 3},
	// {4, 5, 6},
	// {7, 8, 9},
	// i - k <= r <= i + k,
	// j - k <= c <= j + k 且
	// (r, c) 在矩阵内。

	var arr = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// 3 * 3
	// 2 <= r <= 4
	// 2 <= c <= 4
	// mat[]
	// [12,21,16],
	// [27,45,33],
	// [24,39,28]

	// 1 <= r <= 5
	// 1 <= c <= 2
	// [45,45,45],
	// [45,45,45],
	// [45,45,45]
	var target = 2
	fmt.Println(matrixBlockSum(arr, target))
}

//
// matrixBlockSum
// 说实话，愣是没看懂题目表达的是啥意思，现在我又懂了（-_-|）
// 看了下题解，使用二维前缀和解决
//
func matrixBlockSum(mat [][]int, k int) [][]int {
	var yLen = len(mat)
	var xLen = len(mat[0])
	var result = make([][]int, yLen)
	var sum = make([][]int, yLen+1)
	for i := 0; i < yLen; i++ {
		result[i] = make([]int, xLen)
		sum[i] = make([]int, xLen+1)
	}
	sum[yLen] = make([]int, xLen+1)
	for y := 1; y <= yLen; y++ {
		for x := 1; x <= xLen; x++ {
			sum[y][x] = sum[y-1][x] + sum[y][x-1] - sum[y-1][x-1] + mat[y-1][x-1]
		}
	}
	var left, right, top, end = 0, 0, 0, 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if x-k > 0 {
				left = x - k
			} else {
				left = 0
			}
			if x+k < xLen {
				right = x + k
			} else {
				right = xLen - 1
			}
			if y-k > 0 {
				top = y - k
			} else {
				top = 0
			}
			if y+k < yLen {
				end = y + k
			} else {
				end = yLen - 1
			}
			result[y][x] = sum[end+1][right+1] - sum[end+1][left] - sum[top][right+1] + sum[top][left]
		}
	}
	return result
}

//
// matrixBlockSum
// 说实话，愣是没看懂题目表达的是啥意思，现在我又懂了（-_-|）
// 暴力破解
//
func matrixBlockSum1(mat [][]int, k int) [][]int {
	var yLen = len(mat)
	var xLen = len(mat[0])
	var result = make([][]int, yLen)
	for i := 0; i < yLen; i++ {
		result[i] = make([]int, xLen)
	}
	for i := 0; i < yLen; i++ {
		for j := 0; j < xLen; j++ {
			for y := i - k; y <= i+k; y++ {
				for x := j - k; x <= j+k; x++ {
					if y >= 0 && y < yLen && x >= 0 && x < xLen {
						result[y][x] += mat[i][j]
					}
				}
			}
		}
	}
	return result
}
