//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/15 16:54
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	var numMatrix = Constructor(arr)
	var result1 = numMatrix.SumRegion(2, 1, 4, 3) // return 8 (红色矩形框的元素总和)
	var result2 = numMatrix.SumRegion(1, 1, 2, 2) // return 11 (绿色矩形框的元素总和)
	var result3 = numMatrix.SumRegion(1, 2, 2, 4) // return 12 (蓝色矩形框的元素总和)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

type NumMatrix struct {
	xLen int
	yLen int
	sum  [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var instance = NumMatrix{}
	instance.yLen = len(matrix)
	instance.xLen = len(matrix[0])
	var sum = make([][]int, instance.yLen+1)
	for i := 0; i <= instance.yLen; i++ {
		sum[i] = make([]int, instance.xLen+1)
	}
	for y := 1; y <= instance.yLen; y++ {
		for x := 1; x <= instance.xLen; x++ {
			sum[y][x] = sum[y-1][x] + sum[y][x-1] - sum[y-1][x-1] + matrix[y-1][x-1]
		}
	}
	instance.sum = sum
	return instance
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row2+1][col1] - this.sum[row1][col2+1] + this.sum[row1][col1]
}
