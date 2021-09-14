//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/14 10:32
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var target = 6
	fmt.Println(generate(target))
}

func generate(numRows int) [][]int {
	var result = make([][]int, 0)
	var lastArr []int
	for i := 1; i <= numRows; i++ {
		var tmp = make([]int, i)
		tmp[0] = 1
		tmp[len(tmp)-1] = 1
		if len(result) > 0 {
			lastArr = result[i-2]
			for j := 1; j < len(tmp)-1; j++ {
				tmp[j] = lastArr[j-1] + lastArr[j]
			}
		}
		result = append(result, tmp)
	}
	return result
}
