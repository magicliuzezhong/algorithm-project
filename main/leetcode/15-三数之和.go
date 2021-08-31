//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/19 14:44
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{-1, 0, 1, 2, -1, 1, 1, 1, 1}
	//var arr = []int{}
	var result = threeSum(arr)
	fmt.Println(result)
}

//
// threeSum
// 三数之和，第一思路不对，不能用回溯，回溯是有重复值的，说明关键点是在如何处理数据去重的问题
//
func threeSum(nums []int) [][]int {
	var result = make([][]int, 0)
	var numLen = len(nums)
	sort.Ints(nums)
	for first := 0; first < numLen; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		var end = numLen - 1
		for left := first + 1; left < numLen; left++ {
			if left > first+1 && nums[left] == nums[left-1] {
				continue
			}
			for left < end && nums[first]+nums[left]+nums[end] > 0 {
				end--
			}
			if left == end {
				break
			}
			if nums[first]+nums[left]+nums[end] == 0 {
				result = append(result, []int{nums[first], nums[left], nums[end]})
			}
		}
	}
	return result
}
