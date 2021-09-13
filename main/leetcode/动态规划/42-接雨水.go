//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/13 17:25
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var arr = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	var arr = []int{4, 2, 0, 3, 2, 5}
	//var arr = []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}
	fmt.Println(trap(arr))
}

//
// trap
// 先用双指针做出来再想动态规划的解法
//
func trap(height []int) int {
	// 0,1,0,2,1,0,1,3,2,1,2,1
	// l                     r
	// l                     r    low = 0  h = 1  val = 0
	//   l                   r    low = 1  h = 1  val = 0
	//     l                 r    low = 1  h = 1  val = 1
	//       l               r    low = 1  h = 2  val = 1
	//         l           r      low = 1  h = 2  val = 2
	//         l           r      low = 1  h = 2  val = 2

	var heightLen = len(height)
	var value = 0
	var left = 0
	var right = heightLen - 1
	var leftValue = height[left]
	var rightValue = height[right]
	for left <= right {
		if leftValue <= rightValue {
			if height[left] < leftValue {
				value += leftValue - height[left]
			} else {
				leftValue = height[left]
			}
			left++
		} else {
			if height[right] < rightValue {
				value += rightValue - height[right]
			} else {
				rightValue = height[right]
			}
			right--
		}
	}
	return value
}
