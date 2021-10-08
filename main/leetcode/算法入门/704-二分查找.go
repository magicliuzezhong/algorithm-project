//
// Package 算法入门
// @Description：
// @Author：liuzezhong 2021/9/30 22:54
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{-1, 0, 3, 5, 9, 12}
	var result = search(arr, 2)
	fmt.Println(result)
}

func search(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1
	for left < right {
		var mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
