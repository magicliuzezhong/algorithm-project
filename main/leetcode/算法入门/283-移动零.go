//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/9 15:15
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{2, 1, 0, 3, 0, 12, 0}
	moveZeroes(arr)
	fmt.Println(arr)
}

func moveZeroes(nums []int) {
	var left, right, numLen = 0, 0, len(nums)
	for right < numLen {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func moveZeroes1(nums []int) {
	var rightVal = 0
	var numLen = len(nums)
	for i := 0; i < numLen; i++ {
		if nums[i] == 0 {
			if rightVal < i {
				rightVal = i + 1
			}
			for j := rightVal; j < numLen; j++ {
				if nums[j] != 0 {
					rightVal = j
					break
				}
			}
			if rightVal < numLen && i < rightVal {
				nums[i], nums[rightVal] = nums[rightVal], nums[i]
			}
		}
	}
}
