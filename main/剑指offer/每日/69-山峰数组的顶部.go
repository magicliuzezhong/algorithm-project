//
// Package 剑指offer
// @Description：
// @Author：liuzezhong 2021/10/14 10:50
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{2, 3, 4, 5, 6, 3, 2, 1}
	var result = peakIndexInMountainArray(arr)
	fmt.Println(result)
}

// 3, 4, 5, 4, 3, 2, 1
//          1
func peakIndexInMountainArray(arr []int) int {
	var left = 1
	var right = len(arr) - 2
	for left <= right {
		var mid = (right + left) >> 1
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid-1] < arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func peakIndexInMountainArray1(arr []int) int {
	var max = -1
	var maxI = -1
	for i, val := range arr {
		if val > max {
			max = val
			maxI = i
		}
	}
	return maxI
}
