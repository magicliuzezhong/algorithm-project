//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/10 10:47
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{3, 2, 1}
	nextPermutation(arr)
	fmt.Println(arr)
}


func nextPermutation(nums []int) {
	var numLen = len(nums)
	// 从后往前找出第一个不是降序的索引
	var firstIndex = numLen - 1
	for i := numLen - 1; i >= 0; i-- {
		if i-1 >= 0 && nums[i] > nums[i-1] {
			firstIndex = i - 1
			break
		}
	}
	var secondIndex = firstIndex
	for i := numLen - 1; i > firstIndex; i-- {
		if nums[firstIndex] < nums[i] {
			secondIndex = i
			break
		}
	}
	if firstIndex != secondIndex { //交换其位置
		nums[firstIndex], nums[secondIndex] = nums[secondIndex], nums[firstIndex]
	}
	var left = firstIndex + 1
	if left == numLen {
		left = 0
	}
	var right = numLen - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
