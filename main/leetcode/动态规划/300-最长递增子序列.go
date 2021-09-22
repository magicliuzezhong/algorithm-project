//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/22 15:53
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	//var arr = []int{10, 9, 2, 5, 3, 7, 101, 18}
	var arr = []int{0, 1, 0, 3, 2, 3}
	//var arr = []int{0}
	//var arr = []int{7, 7, 7, 7, 7, 7, 7}
	var result = lengthOfLIS(arr)
	fmt.Println(result)
}

//
// lengthOfLIS
// 第一时间想到是这个，动态规划，然后不断往前找最小的哪一个值
//
func lengthOfLIS(nums []int) int {
	var numLen = len(nums)
	var dp = make([]int, numLen)
	dp[0] = 1
	var val = 1
	for i := 1; i < numLen; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > val {
			val = dp[i]
		}
	}
	return val
}
