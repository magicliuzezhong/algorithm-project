//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/21 14:33
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	var digits = []int{}
	var result = plusOne(digits)
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			return digits
		}
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}
