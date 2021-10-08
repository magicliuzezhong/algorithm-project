//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/8 15:52
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	//var s = "  Hello, my name is John"
	var s = " 1 aw"
	var result = countSegments(s)
	fmt.Println(result)
}

func countSegments(s string) int {
	var sLen = len(s)
	var i = 0
	var count = 0
	for i < sLen {
		if s[i] != ' ' {
			count++
		}
		for i < sLen && s[i] != ' ' {
			i++
		}
		i++
	}
	return count
}
