//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/7 14:20
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	var arr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	var result = maxSubArray(arr)
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	var numLen = len(nums)
	var maxVal = nums[0]
	for i := 1; i < numLen; i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		maxVal = maxFunc(maxVal, nums[i])
	}
	return maxVal
	//var sum = nums[0]
	//var maxVal = sum
	//for i := 1; i < numLen; i++ {
	//	sum = maxFunc(sum + nums[i], nums[i])
	//	maxVal = maxFunc(maxVal, sum)
	//}
	//return maxVal
}

func maxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}
