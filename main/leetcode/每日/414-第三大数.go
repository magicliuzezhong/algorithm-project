//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/8 16:07
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	//var arr = []int{3, 4, 1, 2, 6,8}
	var arr = []int{3, 4}
	var result = thirdMax(arr)
	fmt.Println(result)
}

func thirdMax(nums []int) int {
	var maxVal = math.MinInt64
	var midVal = math.MinInt64
	var minVal = math.MinInt64
	for _, num := range nums {
		if num == maxVal || num == midVal || num == minVal {
			continue
		}
		if num > maxVal {
			minVal = midVal
			midVal = maxVal
			maxVal = num
		} else if num > midVal {
			minVal = midVal
			midVal = num
		} else if num > minVal {
			minVal = num
		}
	}
	if minVal != math.MinInt64 {
		return minVal
	} else {
		return maxVal
	}
}
