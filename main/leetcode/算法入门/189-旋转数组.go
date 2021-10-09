//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/9 14:36
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var arr = []int{1, 2, 3, 4, 5, 6, 7}
	//var arr = []int{-1, -100, 3, 99}
	var arr = []int{1, 2, 3}
	rotate(arr, 4)
	fmt.Println(arr)
}

func rotate(nums []int, k int) {
	var numLen = len(nums)
	var mid = numLen - k
	if k > numLen {
		mid = numLen - k%numLen
	}
	var newArr = make([]int, numLen)
	var index = 0
	var right = mid
	for right < numLen {
		newArr[index] = nums[right]
		right++
		index++
	}
	var left = 0
	for left < numLen && left < mid {
		newArr[index] = nums[left]
		left++
		index++
	}
	for i, val := range newArr {
		nums[i] = val
	}
}

func rotate1(nums []int, k int) {
	for j := 0; j < k; j++ {
		var back = nums[0]
		for i := 1; i < len(nums); i++ {
			back, nums[i] = nums[i], back
		}
		nums[0] = back
	}
}
