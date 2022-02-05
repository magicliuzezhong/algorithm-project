// Package main
// @Author 刘泽忠
// @Time 2022-02-05 22:02:16
// @Description: 黄金矿工
package main

import (
	"fmt"
)

func main() {
	//var grid = [][]int{
	//	{1, 0, 7},
	//	{2, 0, 6},
	//	{3, 4, 5},
	//	{0, 3, 0},
	//	{9, 0, 20},
	//}

	var grid = [][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	}
	res1 := getMaximumGold(grid)
	fmt.Println(res1)
}

func getMaximumGold(grid [][]int) int {
	st := make([][]bool, len(grid))
	lens := len(grid[0])
	for i, _ := range grid {
		st[i] = make([]bool, lens)
	}
	res := 0
	var dfs func(y, x, sum int)
	dfs = func(y, x, sum int) {
		if sum > res {
			res = sum
		}
		st[y][x] = true

		dx := [4]int{0, 0, 1, -1}
		dy := [4]int{1, -1, 0, 0}
		for i := 0; i < 4; i++ {
			y := y + dy[i]
			x := x + dx[i]
			if x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) && !st[y][x] && grid[y][x] > 0 {
				dfs(y, x, sum+grid[y][x])
			}
		}
		st[y][x] = false
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 0 {
				dfs(i, j, grid[i][j])
			}
		}
	}
	return res
}
