//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/6 09:32
// @Company cloud-ark.com
//
package main

import "fmt"

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
// 偷窃到的最高金额 = 1 + 3 = 4 。
//
func main() {
	var arr = []int{2, 7, 9, 3, 1}
	//var arr = []int{1, 2, 3, 1}
	var result = rob1(arr)
	fmt.Println(result)
}

func rob1(nums []int) int {
	var numLen = len(nums)
	if numLen == 0 {
		return 0
	}
	// dp[n] = max(dp[n - 2] + arr[n], dp[n-1])
	// 被我一次性就想到了，根本不需要看题解，哈哈
	var dp = make([]int, numLen)
	dp[0] = nums[0]
	for i := 1; i < numLen; i++ {
		if i-2 >= 0 {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		} else {
			dp[i] = max(nums[0], nums[1])
		}
	}
	return dp[numLen-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
