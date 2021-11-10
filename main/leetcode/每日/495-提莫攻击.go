//
// Package main
// @Description：
// @Author：liuzezhong 2021/11/10
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{1}
	var d = 2
	var res = findPoisonedDuration(arr, d)
	fmt.Println(res)
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	var counts = 0
	var arrLen = len(timeSeries)
	for i := 0; i < arrLen; i++ {
		if i+1 < arrLen && timeSeries[i+1]-timeSeries[i] < duration {
			counts += timeSeries[i+1] - timeSeries[i]
		} else {
			counts += duration
		}
	}
	return counts
}
