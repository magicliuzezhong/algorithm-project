//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/12 17:09
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var grid = [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	var res = maxAreaOfIsland(grid)
	fmt.Println(res)
}

func maxAreaOfIsland(grid [][]int) int {
	var yLen, xLen = len(grid), len(grid[0])
	var dfs func(grid [][]int, y, x int, count *int)
	dfs = func(grid [][]int, y, x int, count *int) {
		if y < 0 || x < 0 || y >= yLen || x >= xLen || grid[y][x] == 0 {
			return
		}
		*count++
		grid[y][x] = 0
		dfs(grid, y-1, x, count)
		dfs(grid, y+1, x, count)
		dfs(grid, y, x-1, count)
		dfs(grid, y, x+1, count)
	}
	var res = 0
	for i := 0; i < yLen; i++ {
		for j := 0; j < xLen; j++ {
			if grid[i][j] == 0 {
				continue
			}
			var count = 0
			dfs(grid, i, j, &count)
			if count > res {
				res = count
			}
		}
	}
	return res
}
