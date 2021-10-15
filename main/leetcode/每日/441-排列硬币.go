//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/15 14:21
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n = 5
	var result = arrangeCoins(n)
	fmt.Println(result)
}

func arrangeCoins(n int) int {
	//var left = 1
	//var right = n
	//for left < right {
	//	var mid = (left + right + 1) >> 1
	//	if mid*(mid+1) <= 2*n {
	//		left = mid
	//	} else {
	//		right = mid - 1
	//	}
	//}
	//return left
	return sort.Search(n, func(i int) bool {
		i++
		return i*(i+1) > 2*n
	})
}

func arrangeCoins1(n int) int {
	// 1 + 2 + 3 + 4
	// ans = n + n*(n-1) / 2 = 8
	var shell = 0
	for n > 0 {
		shell++
		n -= shell
		if n < 0 {
			shell--
			break
		}
	}
	return shell
}
