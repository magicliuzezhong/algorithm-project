//
// Package main
// @Description：
// @Author：liuzezhong 2021/11/15
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	//var n = 10000
	var count = make(map[int]int)
	for i := 0; i < 30; i++ {
		var res = bulbSwitch(i)
		count[res]++
		fmt.Println("i: ", i, "res: ", res)
	}
	for key, val := range count {
		fmt.Println("key：", key, "value：", val)
	}
}

// 完全没看题解自主完成(-_-)
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

// 暴力搜索发现这是有规律的
func bulbSwitch1(n int) int {
	var arr = make([]bool, n)
	for index, _ := range arr {
		arr[index] = true
	}
	var turn = 2
	var counts = n
	for turn <= n {
		for i := turn - 1; i < n; i += turn {
			if arr[i] {
				counts--
			} else {
				counts++
			}
			arr[i] = !arr[i]
		}
		//fmt.Println(arr)
		turn++
	}
	return counts
}
