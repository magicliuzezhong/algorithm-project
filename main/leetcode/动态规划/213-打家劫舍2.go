//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/6 09:50
// @Company cloud-ark.com
//
package main

import "fmt"

//
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
// 这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
// 同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
func main() {
	var arr = []int{1, 2, 3, 1}
	//var arr = []int{1, 2, 3, 1, 2}
	var result = rob(arr)
	fmt.Println(result)
}

// 这里是rob那么198的题目改名为rob1，不然冲突了
func rob(nums []int) int {
	var numsLen = len(nums)
	// dp[n] = max(dp[n - 2] + num[n], dp[n - 1])
	if numsLen == 0 {
		return 0
	} else if numsLen == 1 {
		return nums[0]
	} else if numsLen == 2 {
		return maxRange(nums[0], nums[1])
	}
	return maxRange(robRange(nums, 0, numsLen-2), robRange(nums, 1, numsLen-1))
}

func robRange(nums []int, left, right int) int {
	var numsLen = len(nums)
	var dp = make([]int, numsLen)
	dp[left] = nums[left]
	for i := left + 1; i <= right; i++ {
		if i-2 >= left {
			dp[i] = maxRange(dp[i-2]+nums[i], dp[i-1])
		} else {
			dp[i] = maxRange(nums[left], nums[i])
		}
	}
	return dp[right]
}

func maxRange(a, b int) int {
	if a > b {
		return a
	}
	return b
}
