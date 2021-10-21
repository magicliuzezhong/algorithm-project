//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/21 14:50
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n = 4
	var result = solveNQueens(n)
	fmt.Println(result)
}

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	var queue = make([]int, n)
	var columns = make(map[int]bool)
	var left = make(map[int]bool)
	var right = make(map[int]bool)
	backtrack(queue, n, 0, columns, left, right)
	return res
}

//
// backtrack
// 非常经典的回溯算法，N皇后问题
//
func backtrack(queens []int, n, row int, columns, left, right map[int]bool) {
	if row == n {
		var tmp = make([]string, n)
		for index, val := range queens {
			var sb = strings.Builder{}
			for i := 0; i < n; i++ {
				if i == val {
					sb.WriteString("Q")
				} else {
					sb.WriteString(".")
				}
			}
			tmp[index] = sb.String()
		}
		res = append(res, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] { // 列已经存在，返回
			continue
		}
		var leftCount = row + i
		if left[leftCount] { //左斜已经存在
			continue
		}
		var rightCount = row - i
		if right[rightCount] {
			continue
		}
		// 可以放置数据
		left[leftCount] = true
		right[rightCount] = true
		columns[i] = true
		queens[row] = i
		backtrack(queens, n, row+1, columns, left, right)
		delete(left, leftCount)
		delete(right, rightCount)
		delete(columns, i)
		queens[row] = -1
	}
}
