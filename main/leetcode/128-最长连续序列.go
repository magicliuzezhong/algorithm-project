//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/19 11:21
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	//var arr = []int{100, 4, 200, 1, 3, 2}
	var arr = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	var count = longestConsecutive(arr)
	fmt.Println(count)
}

func longestConsecutive(nums []int) int {
	var record = map[int]struct{}{}
	for _, val := range nums {
		record[val] = struct{}{}
	}
	var maxLong = 0
	var count = 1
	var exists1 = false
	var exists2 = false
	for _, val := range nums {
		count = 1
		if _, exists1 = record[val-1]; !exists1 {
			for {
				if _, exists2 = record[val+count]; exists2 {
					count++
				} else {
					break
				}
			}
		}
		if count > maxLong {
			maxLong = count
		}
	}
	return maxLong
}
