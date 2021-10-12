//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/12 15:11
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 16
	var b = 5
	var res = divide(a, b)
	fmt.Println(res)
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}
	var a = dividend
	var b = divisor
	var sign = 1
	if a > 0 && b < 0 || a < 0 && b > 0 {
		sign = -1
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	var res = div(int64(a), int64(b))
	if sign > 0 {
		return res
	}
	return -res
}

func div(a, b int64) int {
	if a < b {
		return 0
	}
	var count = 1
	var tb = b
	for (tb + tb) <= a {
		count = count << 1
		tb = tb << 1
	}
	return count + div(a-tb, b)
}
