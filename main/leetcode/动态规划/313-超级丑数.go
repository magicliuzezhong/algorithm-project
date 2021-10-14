//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/14 14:44
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	var n = 12
	var primes = []int{2, 7, 13, 19}
	var result = nthSuperUglyNumber(n, primes)
	fmt.Println(result)
}

func nthSuperUglyNumber(n int, primes []int) int {
	var primeLen = len(primes)
	var site = make([]int, primeLen)
	var nums = make([]int, primeLen)
	for i, _ := range site {
		site[i] = 1
	}
	var dp = make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < primeLen; j++ {
			nums[j] = dp[site[j]] * primes[j]
		}
		dp[i] = minNums(nums...)
		for j := 0; j < primeLen; j++ {
			if dp[i] == nums[j] {
				site[j]++
			}
		}
	}
	return dp[n]
}

func minNums(nums ...int) int {
	var minVal = math.MaxInt64
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}
