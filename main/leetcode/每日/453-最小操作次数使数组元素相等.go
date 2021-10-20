//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/20 22:03
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 1000000000}
	var result = minMoves(arr)
	fmt.Println(result)
}

// 坏了，今天的简单题都不能重拳出击了（-_-）
func minMoves(nums []int) int {
	var count = 0
	var minVal = nums[0]
	for _, val := range nums[1:] {
		if val < minVal {
			minVal = val
		}
	}
	for _, val := range nums {
		count += val - minVal
	}
	return count
}
