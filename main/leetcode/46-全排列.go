//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/18 10:56
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4}
	var result = permute(arr)
	fmt.Println(result)
}

func permute(nums []int) [][]int {
	var result = make([][]int, 0)
	var tmpResult = make([]int, 0)
	listPermute(nums, -1, &result, &tmpResult)
	return result
}

func listPermute(nums []int, start int, result *[][]int, tmpResult *[]int) {
	var numLen = len(nums)
	if len(*tmpResult) == numLen {
		var newResult = make([]int, numLen)
		for i, val := range *tmpResult {
			newResult[i] = val
		}
		*result = append(*result, newResult)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i == start {
			continue
		}
		var basic = nums[i]
		var exists = false
		for _, tmpVal := range *tmpResult {
			if basic == tmpVal {
				exists = true
				break
			}
		}
		if exists {
			continue
		}
 		*tmpResult = append(*tmpResult, basic)
		listPermute(nums, i, result, tmpResult)
		var tmpLen = len(*tmpResult)
		if tmpLen > 0 {
			*tmpResult = append((*tmpResult)[:tmpLen - 1], (*tmpResult)[tmpLen:]...)
		}
	}
}
