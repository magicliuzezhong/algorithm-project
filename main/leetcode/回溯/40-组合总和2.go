//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/26 16:47
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"sort"
)

func main() {
	//var arr = []int{10, 1, 2, 7, 6, 1, 5}
	var arr = []int{1, 1, 2, 5, 6, 7, 10}
	var target = 8
	var res = combinationSum2(arr, target)
	fmt.Println(res)
}

//
// combinationSum2
// 这也是一道非常经典的回溯算法
//
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result = make([][]int, 0)
	var tmp = make([]int, 0)
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if target < 0 {
			return
		}
		if target == 0 {
			result = append(result, append([]int{}, tmp...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i-1] == candidates[i] {
				continue
			}
			tmp = append(tmp, candidates[i])
			dfs(target-candidates[i], i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(target, 0)
	return result
}
