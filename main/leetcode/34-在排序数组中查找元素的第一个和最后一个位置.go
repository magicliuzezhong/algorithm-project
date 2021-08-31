//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/23 10:29
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	//var arr = []int{1, 2, 3, 4, 5, 5, 7, 7, 7, 8, 8, 8, 8, 9, 10}
	var arr = []int{-3,-2,-1}
	var ranges = searchRange(arr, 4)
	fmt.Println(ranges)
}

//
// searchRange
// 要求算法复杂度 O(log n)  所以必须要使用二分查找
//
func searchRange(nums []int, target int) []int {
	var left = binarySearch(nums, target, true)
	var right = binarySearch(nums, target, false)
	return []int{left, right}
}

func binarySearch(nums []int, target int, isLeft bool) int {
	var numLen = len(nums)
	var left = 0
	var right = numLen - 1
	var mid = 0
	var ans = -1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			ans = mid
			if isLeft { //需要找到最左元素
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return ans
}

//
// searchRange1
// 算法复杂度 O(n)
//
func searchRange1(nums []int, target int) []int {
	var numLen = len(nums)
	if numLen == 1 && nums[0] != target {
		return []int{-1, -1}
	}
	var left = 0
	var right = numLen - 1
	for left <= right {
		if nums[left] != target {
			left++
		}
		if nums[right] != target {
			right--
		}
		if left <= right && nums[left] == target && nums[right] == target {
			break
		}
	}
	if left > right {
		return []int{-1, -1}
	}
	return []int{left, right}
}
