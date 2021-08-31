//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/5 17:31
// @Company cloud-ark.com
//
package main

import "fmt"

//假设你正在爬楼梯。需要 n阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
//
//示例 1：
//
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//示例 2：
//
//输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
func main() {
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	var df = make([]int, n+1)
	df[1] = 1
	df[2] = 2
	for i := 3; i <= n; i++ {
		df[i] = df[i-1] + df[i-2]
	}
	return df[n]
}
