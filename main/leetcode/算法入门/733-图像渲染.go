//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/12 16:44
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var image = [][]int{
		{1, 1, 1},
		{1, 1, 3},
		{1, 3, 3},
	}
	var res = floodFill(image, 2, 1, 2)
	for _, re := range res {
		fmt.Println(re)
	}
}

// 1 1 1    2 2 2
// 1 1 0    2 2 0
// 1 0 1    2 0 1

// 1 1
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var yLen = len(image)
	var xLen = len(image[0])
	var dfs func(image [][]int, y, x, color, newColor int)
	dfs = func(image [][]int, y, x, color, newColor int) {
		if image[y][x] == newColor {
			return
		}
		if image[y][x] != color {
			return
		}
		image[y][x] = newColor
		if x-1 >= 0 {
			dfs(image, y, x-1, color, newColor)
		}
		if x+1 < xLen {
			dfs(image, y, x+1, color, newColor)
		}
		if y-1 >= 0 {
			dfs(image, y-1, x, color, newColor)
		}
		if y+1 < yLen {
			dfs(image, y+1, x, color, newColor)
		}
	}
	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}
