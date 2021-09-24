//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/24 13:58
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var arr = []int{1, 3, 5, 4, 7}
	//var arr = []int{1, 2, 4, 3, 5, 4, 7, 2}
	var arr = []int{2, 2, 2, 2, 2}
	var result = findNumberOfLIS(arr)
	fmt.Println(result)
}

func findNumberOfLIS(nums []int) int {
	var numLen = len(nums)
	var dp = make([]int, numLen)
	var dpRecord = make([]int, numLen)
	var val = 0
	var valCount = 0
	for i := 0; i < numLen; i++ {
		dp[i] = 1
		dpRecord[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					dpRecord[i] = dpRecord[j]
				} else if dp[j]+1 == dp[i] {
					dpRecord[i] += dpRecord[j]
				}
			}
		}
		if val < dp[i] {
			val = dp[i]
			valCount = dpRecord[i]
		} else if val == dp[i] {
			valCount += dpRecord[i]
		}
	}
	return valCount
}
