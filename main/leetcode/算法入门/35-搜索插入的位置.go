//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/9 10:59
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{1, 3, 5, 6}
	var result = searchInsert(arr, 0)
	fmt.Println(result)
}

func searchInsert(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1
	var mid = 0
	for left <= right {
		mid = (left + right) / 2
		if mid == 0 && nums[mid] > target {
			return 0
		} else if mid == len(nums)-1 && nums[mid] < target {
			return mid + 1
		} else if nums[mid] == target {
			return mid
		} else if mid > 0 && nums[mid-1] < target && nums[mid] > target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}
