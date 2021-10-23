//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/23 12:03
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	var area = 15
	var result = constructRectangle(area)
	fmt.Println(result)
}

func constructRectangle(area int) []int {
	var w = int(math.Sqrt(float64(area)))
	for area%w > 0 {
		w--
	}
	return []int{area / w, w}
}

func constructRectangle1(area int) []int {
	var left = 1
	var right = area
	var expectLeft = 1
	var expectRight = area
	for left <= right {
		var val = left * right
		if val > area {
			right--
		} else if val == area {
			expectLeft = left
			expectRight = right
			left++
			right--
		} else {
			left++
		}
	}
	return []int{expectLeft, expectRight}
}
