//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/18 08:58
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

//编写一个程序，通过填充空格来解决数独问题。
//
//数独的解法需 遵循如下规则：
//
//数字1-9在每一行只能出现一次。
//数字1-9在每一列只能出现一次。
//数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
//数独部分空格内已填入了数字，空白格用'.'表示。
//
func main() {
	var board = [][]byte{
		//0 1  2  3  4  5  6  7  8
		{5, 3, 0, 0, 7, 0, 0, 0, 0}, // 0
		{6, 0, 0, 1, 9, 5, 0, 0, 0}, // 1
		{0, 9, 8, 0, 0, 0, 0, 6, 0}, // 2

		{8, 0, 0, 0, 6, 0, 0, 0, 3}, // 3
		{4, 0, 0, 8, 0, 3, 0, 0, 1}, // 4
		{7, 0, 0, 0, 2, 0, 0, 0, 6}, // 5

		{0, 6, 0, 0, 0, 0, 2, 8, 0}, // 6
		{0, 0, 0, 4, 1, 9, 0, 0, 5}, // 7
		{0, 0, 0, 0, 8, 0, 0, 7, 9}, // 8
	}

	solveSudoku(board)
	for _, val := range board {
		fmt.Println(val)
	}
	//var aa = expected3x3(board, 2, 3)

	//fmt.Println(aa)
	//fmt.Println(board[1][4])
}

func solveSudoku(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				var expectedVal = expected3x3(board, i, j)
				var expectedH = checkH(board, i, j)
				var expectedV = checkV(board, i, j)
				var expectedData = mergeExpected(expectedVal, expectedH, expectedV)
				if len(expectedData) == 1 {
					for key, _ := range expectedData {
						board[i][j] = key
					}
					solveSudoku(board)
				}
			}
		}
	}
}

func mergeExpected(expectedVal, expectedH, expectedV map[byte]bool) map[byte]bool {
	var record = make(map[byte]bool)
	for key, _ := range expectedVal {
		_, ok1 := expectedH[key]
		_, ok2 := expectedV[key]
		if ok1 && ok2 {
			record[key] = true
		}
	}
	return record
}

func checkV(board [][]byte, i, j int) map[byte]bool {
	var record = make(map[byte]bool)
	if board[i][j] != 0 {
		return record
	}
	for val := byte(1); val < 10; val++ {
		record[val] = true
	}
	for v := 0; v < 9; v++ {
		if board[i][v] != 0 {
			delete(record, board[i][v])
		}
	}
	return record
}

func checkH(board [][]byte, i, j int) map[byte]bool {
	var record = make(map[byte]bool)
	if board[i][j] != 0 {
		return record
	}
	for val := byte(1); val < 10; val++ {
		record[val] = true
	}
	for h := 0; h < 9; h++ {
		if board[h][j] != 0 {
			delete(record, board[h][j])
		}
	}
	return record
}

func expected3x3(board [][]byte, i, j int) map[byte]bool {
	var record = make(map[byte]bool)
	if board[i][j] != 0 {
		return record
	}
	for val := byte(1); val < 10; val++ {
		record[val] = true
	}
	var x = (i / 3) * 3
	var y = (j / 3) * 3
	for h := x; h < x+3; h++ {
		for v := y; v < y+3; v++ {
			if int(board[h][v]) != 0 {
				delete(record, board[h][v])
			}
		}
	}
	return record
}
