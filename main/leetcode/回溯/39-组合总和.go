//
// Package 回溯
// @Description：
// @Author：liuzezhong 2021/9/26 16:08
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{2, 3, 6, 7}
	var target = 7
	var result = combinationSum(arr, target)
	fmt.Println(result)
}

//
// combinationSum
// 非常经典的回溯算法，剪枝非常得经典
//
func combinationSum(candidates []int, target int) [][]int {
	var tmp = make([]int, 0)
	var result = make([][]int, 0)
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			result = append(result, append([]int{}, tmp...))
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			tmp = append(tmp, candidates[idx])
			dfs(target-candidates[idx], idx)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(target, 0)
	return result
}
