//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/14 09:03
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 8, 9, 10, 11, 12, 13}
	fmt.Println(numberOfArithmeticSlices(arr))
}

//
// numberOfArithmeticSlices
// 最开始使用的dp
// 发现dp只与倒数第二项有关，使用了pre，然后优化了空间
// 没有看题解（-_-）
//
func numberOfArithmeticSlices(nums []int) int {
	// 3  1
	// 4  3
	// 5  6
	// 6  10  4 + 3 + 2 + 1
	// 获取等差数列最长的值
	// dp[i] = dp[i - 2] + nums[i] - nums[i - 1] == nums[i - 1] - nums[i - 2]
	var numLen = len(nums)
	if numLen < 2 {
		return 0
	}
	var value = 0
	var pre = 0
	for i := 2; i < numLen; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			if pre == 0 {
				pre = 3
			} else {
				pre = pre + 1
			}
		} else {
			value += calcShell(pre)
			pre = 0
		}
	}
	value += calcShell(pre)
	return value
}

func calcShell(num int) int {
	var val = 0
	num -= 2
	for num > 0 {
		val += num
		num--
	}
	return val
}
