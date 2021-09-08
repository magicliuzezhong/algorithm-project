//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/7 16:01
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{1, -2, 3, -2, 1, 5}
	var result = maxSubarraySumCircular(arr)
	var result1 = maxSubarraySumCircular1(arr)
	fmt.Println(result)
	fmt.Println(result1)
}

// 3, -1, 2, -1
// 3   2  4   3

// 看了半天看懂了
// 说白了就是2部分
// 1、最大值在中间  获取maxVal就行
// 2、最大值在两边  使用 sum - minVal（总和 - 最小值）
func maxSubarraySumCircular(nums []int) int {
	var numLen = len(nums)
	var sum = nums[0]
	var maxVal = nums[0]
	var pre = nums[0]
	for i := 1; i < numLen; i++ {
		if pre > 0 {
			pre += nums[i]
		} else {
			pre = nums[i]
		}
		maxVal = maxFunc1(maxVal, pre)
		sum += nums[i]
	}
	var minVal = 0
	pre = nums[0]
	for i := 1; i < numLen-1; i++ {
		if pre < 0 {
			pre += nums[i]
		} else {
			pre = nums[i]
		}
		minVal = minFunc1(minVal, pre)
	}
	return maxFunc1(maxVal, sum-minVal)
}

//
// maxSubarraySumCircular1
// 该方案是上面方案的优化版本
// 优化点：优化了一次循环，两次循环可以合并在一起
//
func maxSubarraySumCircular1(nums []int) int {
	var numLen = len(nums)
	var sum = nums[0]
	var preMax = nums[0]
	var maxVal = nums[0]
	var preMin = 0
	var minVal = 0
	for i := 1; i < numLen; i++ {
		preMax = maxFunc1(preMax+nums[i], nums[i])
		maxVal = maxFunc1(maxVal, preMax)
		if i != numLen-1 {
			preMin = minFunc1(preMin+nums[i], nums[i])
			minVal = minFunc1(minVal, preMin)
		}
		sum += nums[i]
	}
	return maxFunc1(maxVal, sum-minVal)
}

func maxFunc1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minFunc1(a, b int) int {
	if a > b {
		return b
	}
	return a
}
