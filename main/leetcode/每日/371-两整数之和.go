//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/26 15:26
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	var a = math.MaxInt64
	var b = -1
	var result = getSum(a, b)
	fmt.Println(result)
}

// 内存占用 1924 kb
func getSum(a int, b int) int {
	// 0010 ^ 0011 = 0001
	// 0010 & 0011 = 0010
	// 0001 ^ 0010 = 0011
	for b != 0 {
		var addSite = uint(a&b) << 1
		a = a ^ b
		b = int(addSite)
	}
	return a
}

// 内存占用 1928 kb
func getSum1(a int, b int) int {
	var addSite uint = 0
	for b != 0 {
		addSite = uint(a&b) << 1
		a = a ^ b
		b = int(addSite)
	}
	return a
}
