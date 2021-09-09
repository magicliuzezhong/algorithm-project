//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/9 11:25
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{8, 1, 5, 2, 6}

	// score = value[i] + value[j] + i - j
	// score = (value[i] + i) + (value[j] - j)
	// 想要分数高，分数要高，距离要短
	var result = maxScoreSightseeingPair(arr)
	fmt.Println(result)
}

//
// maxScoreSightseeingPair
// 看来下题解，根本看不出来是动态规划问题嘛，完全就是智力题，瞬间感觉自己智力不太行的样子(-_-)
// score = value[i] + value[j] + i - j
// score = (value[i] + i) + (value[j] - j)
//
func maxScoreSightseeingPair(values []int) int {
	var numLen = len(values)
	var maxVal = 0
	var maxI = values[0] + 0
	for i := 1; i < numLen; i++ {
		maxVal = maxFunc4(maxVal, maxI+values[i]-i) //前一个和后一个相加
		maxI = maxFunc4(maxI, values[i]+i)
	}
	return maxVal
}

//
// maxScoreSightseeingPair
// 暴力破解超时，说明复杂度O(n2)是过不了的
//
func maxScoreSightseeingPair1(values []int) int {
	var numLen = len(values)
	var maxVal = 0
	for i := 0; i < numLen; i++ {
		for j := i + 1; j < numLen; j++ {
			maxVal = maxFunc4(maxVal, values[i]+values[j]+i-j)
		}
	}
	return maxVal
}

func maxFunc4(a, b int) int {
	if a > b {
		return a
	}
	return b
}
