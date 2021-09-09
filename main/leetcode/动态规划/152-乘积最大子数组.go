//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/8 14:54
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	//var arr = []int{2, 3, -2, 4}
	//var arr = []int{-2,0,-1}
	var arr = []int{-2, 5, 6, 0, -1, -4}
	fmt.Println(maxProduct(arr))
}

//
// maxProduct
// 记录最大值和最小值，如果nums[i] < 0，那么交换最大值和最大值
//
func maxProduct(nums []int) int {
	var numLen = len(nums)
	var maxVal = math.MinInt64
	var preMax = 1
	var preMin = 1
	for i := 0; i < numLen; i++ {
		if nums[i] < 0 {
			preMax, preMin = preMin, preMax
		}
		preMax = maxFunc2(preMax*nums[i], nums[i])
		preMin = minFunc2(preMin*nums[i], nums[i])

		maxVal = maxFunc2(maxVal, preMax)
	}
	return maxVal
}

func maxFunc2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minFunc2(a, b int) int {
	if a > b {
		return b
	}
	return a
}
