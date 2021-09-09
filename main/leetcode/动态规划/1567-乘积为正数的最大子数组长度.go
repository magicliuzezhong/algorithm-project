//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/8 15:52
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	//var arr = []int{1, -2, -3, 4}
	//var arr = []int{0,1,-2,-3,-4}
	var arr = []int{1, 0, 0, 0, 0}
	var result = getMaxLen(arr)
	fmt.Println(result)
}

func getMaxLen(nums []int) int {
	var numLen = len(nums)
	var maxVal = 0
	var negative = 0
	var positive = 0
	var tmp = 0
	if nums[0] > 0 {
		positive = 1
		maxVal = 1
	} else if nums[0] < 0 {
		negative = 1
	}
	for i := 1; i < numLen; i++ {
		if nums[i] == 0 {
			positive = 0
			negative = 0
		} else if nums[i] > 0 {
			positive++
			if negative > 0 {
				negative++
			}
		} else {
			tmp = negative
			negative = positive + 1
			if tmp > 0 {
				positive = tmp + 1
			} else {
				positive = 0
			}
		}
		maxVal = maxFunc3(maxVal, positive)
	}
	return maxVal
}

func maxFunc3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
