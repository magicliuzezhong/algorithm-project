//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/2 09:41
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var n = 1
	var result = fib(n)
	fmt.Println(result)
}

//
// fib
// f(0) = 0
// f(1) = 1
// f(n) = f(n - 1) + f(n - 2)
// 求f(n)
// 很容易
//
func fib(n int) int {
	if n < 2 {
		return n
	}
	var dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
