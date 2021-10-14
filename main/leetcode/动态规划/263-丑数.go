//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/14 14:49
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var result = isUgly(25)
	fmt.Println(result)
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, val := range []int{2, 3, 5} {
		for n%val == 0 {
			n /= val
		}
	}
	return n == 1
}
