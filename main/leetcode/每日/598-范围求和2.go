package main

import (
	"fmt"
)

func main() {
	var y, x = 3, 3
	var ops = [][]int{
		{2, 2},
		{3, 3},
	}
	var res = maxCount(y, x, ops)
	fmt.Println(res)
}

// EDG是冠军，每日一题庆祝一下
func maxCount(m int, n int, ops [][]int) int {
	var minY = m
	var minX = n
	for _, val := range ops {
		var y, x = val[0], val[1]
		if y < minY {
			minY = y
		}
		if x < minX {
			minX = x
		}
	}
	return minY * minX
}
