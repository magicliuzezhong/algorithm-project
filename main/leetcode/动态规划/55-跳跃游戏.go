//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/6 18:01
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4}
	var result = canJump(arr)
	fmt.Println(result)
}

//
// canJump
// @Description: 贪心算法
// 算法逻辑，
//
func canJump(nums []int) bool {
	var numLen = len(nums)
	var most = 0
	for i := 0; i < numLen-1 && most < numLen-1; i++ {
		most = maxCanJump(most, i+nums[i])
		if most == i {
			break
		}
	}
	return most >= numLen-1
}

func maxCanJump(a, b int) int {
	if a > b {
		return a
	}
	return b
}
