//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/6 10:32
// @Company cloud-ark.com
//
package main

import "fmt"

//
// 给你一个整数数组nums，你可以对它进行一些操作。
//
// 每次操作中，选择任意一个nums[i]，删除它并获得nums[i]的点数。之后，你必须删除 所有 等于nums[i] - 1 和 nums[i] + 1的元素。
//
// 开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
//
func main() {
	// 2
	// 2
	// 3
	// 3
	// 3
	// 4

	// 2 4
	// 3 9
	// 4 4
	// dp[n] = max(dp[n - 2] + hash(n), dp[n - 1])
	//
	//var arr = []int{3, 4, 2}
	var arr = []int{2, 2, 3, 3, 3, 4}
	var result = deleteAndEarn(arr)
	fmt.Println(result)
}

//
// deleteAndEarn
// @Description: hash + 打家劫舍的思路
// 进行hash之后就意味着相邻的元素是不能选取的了，那么按照打家劫舍的思路来处理即可
//
func deleteAndEarn(nums []int) int {
	var maxVal = 0
	for i := 0; i < len(nums); i++ {
		maxVal = maxEarn(nums[i], maxVal)
	}
	var sums = make([]int, maxVal+1)
	for i := 0; i < len(nums); i++ {
		sums[nums[i]] += nums[i]
	}
	return robEarn(sums)
}

func robEarn(sums []int) int {
	var sumLen = len(sums)
	var dp = make([]int, sumLen)
	dp[0] = sums[0]
	for i := 1; i < sumLen; i++ {
		if i >= 2 {
			dp[i] = maxEarn(dp[i-2]+sums[i], dp[i-1])
		} else {
			dp[i] = maxEarn(sums[0], sums[1])
		}
	}
	return dp[sumLen-1]
}

func maxEarn(a, b int) int {
	if a > b {
		return a
	}
	return b
}
