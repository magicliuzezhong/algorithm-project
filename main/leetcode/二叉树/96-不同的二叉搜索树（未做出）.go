//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/24 16:33
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var result = numTrees(19)
	fmt.Println(result)
}

func numTrees(n int) int {
	return helperNumTree(1, n)
}

func helperNumTree(start, end int) int {
	if start > end {
		return 1
	}
	var result = 0
	for i := start; i <= end; i++ {
		var leftCount = helperNumTree(start, i - 1)
		var rightCount = helperNumTree(i+1, end)
		for j := 0; j < leftCount; j++ {
			for k := 0; k < rightCount; k++ {
				result++
			}
		}
	}
	return result
}
