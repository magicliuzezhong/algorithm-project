//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/14 11:46
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var target = 5
	fmt.Println(getRow(target))
}

//
// getRow
// 代码优化，空间复杂度太大，计算i的时候只与i-1相关
//
func getRow(rowIndex int) []int {
	var index, val = 0, 0
	var pre, current = make([]int, rowIndex+1), make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		current[0] = 1
		current[i] = 1
		if i > 1 {
			for j := 1; j < i; j++ {
				current[j] = pre[j-1] + pre[j]
			}
		}
		for index, val = range current {
			pre[index] = val
		}
	}
	return current
}

func getRow1(rowIndex int) []int {
	var result = make([][]int, 0)
	for i := 0; i <= rowIndex; i++ {
		var tmp = make([]int, i+1)
		tmp[0] = 1
		tmp[i] = 1
		if len(result) > 1 {
			var lastArr = result[i-1]
			for j := 1; j < i; j++ {
				tmp[j] = lastArr[j-1] + lastArr[j]
			}
		}
		result = append(result, tmp)
	}
	return result[rowIndex]
}
