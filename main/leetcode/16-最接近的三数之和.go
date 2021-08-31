//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/11 10:07
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
	"sort"
)

// 16. 最接近的三数之和
// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
//示例：
//
// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
func main() {
	var arr = []int{1,1,1,1}
	var val = threeSumClosest(arr, 3)
	fmt.Println(val)
}

// -1,2,1,-4   2
func threeSumClosest(nums []int, target int) int {
	var numLen = len(nums)
	sort.Ints(nums)
	var ans = nums[0] + nums[1] + nums[2]
	for i := 0; i < numLen; i++ {
		var left = i + 1
		var right = numLen - 1
		for left < right {
			var sum = nums[i] + nums[left] + nums[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(target-ans)) {
				ans = sum
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return target
			}
		}
	}
	return ans
}

// 暴力破解
func threeSumClosest1(nums []int, target int) int {
	var numLen = len(nums)

	var minVal = math.MaxInt64
	var result = 0
	var currentVal = 0
	for i := 0; i < numLen; i++ {
		currentVal = nums[i]
		for j := i + 1; j < numLen; j++ {
			currentVal += nums[j]
			for k := j + 1; k < numLen; k++ {
				currentVal += nums[k]
				if int(math.Abs(float64(currentVal-target))) < minVal {
					minVal = int(math.Abs(float64(currentVal - target)))
					result = currentVal
				}
				currentVal -= nums[k]
			}
			currentVal -= nums[j]
		}
	}
	return result
}
