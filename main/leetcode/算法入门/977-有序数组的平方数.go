//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/9 11:13
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var arr = []int{-5, -4, -2, 0, 1, 3, 4, 5, 10}
	var arr = []int{-5, -3}
	var result = sortedSquares(arr)
	fmt.Println(result)
}

// 25  16  4  0  1  9  100
// 9  16  4  0  1  25  100
// 1  16  4  0  9  25  100
// 1  9   4  0  16 25  100
func sortedSquares(nums []int) []int {
	var negativeIndex = -1
	var numLen = len(nums)
	for i := 0; i < numLen; i++ {
		if nums[i] < 0 {
			negativeIndex = i
		} else {
			break
		}
	}
	for i, val := range nums {
		nums[i] = val * val
	}
	var newArr = make([]int, numLen)
	var left = negativeIndex
	var right = negativeIndex + 1
	var index = 0
	for left >= 0 && right < numLen {
		if nums[left] > nums[right] {
			newArr[index] = nums[right]
			right++
		} else {
			newArr[index] = nums[left]
			left--
		}
		index++
	}
	for left >= 0 {
		newArr[index] = nums[left]
		left--
		index++
	}
	for right < numLen {
		newArr[index] = nums[right]
		right++
		index++
	}
	return newArr
}
