package main

import (
	"fmt"
	"strconv"
)

func main() {
	var board = [][]byte{
		//'.' 1  2  3  4  5  6  7  8
		{5, 3, '.', '.', 7, '.', '.', '.', '.'}, // '.'
		{6, '.', '.', 1, 9, 5, '.', '.', '.'},   // 1
		{'.', 9, 8, '.', '.', '.', '.', 6, '.'}, // 2

		{8, '.', '.', '.', 6, '.', '.', '.', 3}, // 3
		{4, '.', '.', 8, '.', 3, '.', '.', 1},   // 4
		{7, '.', '.', '.', 2, '.', '.', '.', 6}, // 5

		{'.', 6, '.', '.', '.', '.', 2, 8, '.'}, // 6
		{'.', '.', '.', 4, 1, 9, '.', '.', 5},   // 7
		{'.', '.', '.', '.', 8, '.', '.', 7, 9}, // 8
	}
	var isValid = isValidSudoku(board)
	fmt.Println(isValid)
}

//
// isValidSudoku
// 有效数独，
// 首先思路，
// 验证横向的所有内容是否包含重复值
// 验证9*9是否有重复值
// 暴力破解我做出来了，就是运行时间有点长
// 看了下官方题解，一个字，6
// var row  = [][]bool{}
// row[0][7] = true，意味着第0行，的7数值已经存在了，下次再次遍历到该数值的时候已经存在了，那么直接不合理
// 一下的两个也是同样的道理
// var col  = [][]bool{}
// box 的第一个索引位要怎么求解？
// var boxIndex = (i / 3) * 3 + j / 3
// var box  = [][]bool{}
//
func isValidSudoku(board [][]byte) bool {
	var row = [9][9]bool{} //记录行内有没有存在
	var col = [9][9]bool{} //记录列内有没有存在
	var box = [9][9]bool{} //记录3*3盒子内有没有存在
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				var basic, _ = strconv.Atoi(string(board[i][j]))
				basic -= 1
				var boxIndex = (i/3)*3 + j/3
				if row[i][basic] || col[j][basic] || box[boxIndex][basic] {
					return false
				}
				row[i][basic] = true
				col[j][basic] = true
				box[boxIndex][basic] = true
			}
		}
	}
	return true
}

func check3x3Right(board [][]byte, i, j int) bool {
	var basicX = (i / 3) * 3
	var basicY = (j / 3) * 3
	var basic = board[i][j]
	for x := basicX; x < basicX+3; x++ {
		for y := basicY; y < basicY+3; y++ {
			if x == i && y == j {
				continue
			}
			if basic == board[x][y] {
				return true
			}
		}
	}
	return false
}

func checkHRight(board [][]byte, i, j int) bool {
	var basic = board[i][j]
	for x := j + 1; x < 9; x++ {
		if basic == board[i][x] {
			return true
		}
	}
	return false
}

func checkVRight(board [][]byte, i, j int) bool {
	var basic = board[i][j]
	for v := i + 1; v < 9; v++ {
		if basic == board[v][j] {
			return true
		}
	}
	return false
}
