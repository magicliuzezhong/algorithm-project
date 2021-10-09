//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/9 16:01
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{2, 3, 4}
	var result = twoSum(arr, 6)
	fmt.Println(result)
}

func twoSum(numbers []int, target int) []int {
	var left = 0
	var right = len(numbers) - 1
	for left < right {
		var res = numbers[left] + numbers[right]
		if res == target {
			return []int{left + 1, right + 1}
		} else if res > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
