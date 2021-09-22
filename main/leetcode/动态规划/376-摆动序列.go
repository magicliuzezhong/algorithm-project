//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/22 17:07
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var arr = []int{1, 7, 4, 9, 2, 5}
	//var arr = []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var arr = []int{84}
	var result = wiggleMaxLength(arr)
	fmt.Println(result)
}

func wiggleMaxLength(nums []int) int {
	// dp[i] = (nums[i] - nums[i-1]) * (nums[i-1] - nums[i-2]) < 0
	// 1  17  5  10  13  15  10       5    16      8
	// 1  2   3  4   4   4   4 / 5    5    4 / 6   4 / 5 / 6
	// 0  +   -  +   +   +   +   -    -    +   +   +   -   +
	var numLen = len(nums)
	var dp = make([]int, numLen)
	var dpRecord = make([]int8, numLen) //记录正负关系
	var currentType int8 = 0
	dp[0] = 1
	var val = 1
	for i := 1; i < numLen; i++ {
		currentType = getType(nums[i], nums[i-1])
		if nums[i] == nums[i-1] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = 2
			dpRecord[i] = currentType
		}
		// 开始规划规则
		for j := 0; j < i; j++ {
			currentType = getType(nums[i], nums[j])
			if dpRecord[j]*currentType < 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				dpRecord[i] = currentType
			}
		}
		if val < dp[i] {
			val = dp[i]
		}
	}
	return val
}

func getType(a, b int) int8 {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
