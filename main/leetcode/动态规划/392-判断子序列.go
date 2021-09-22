//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/22 17:52
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var s = ""
	var t = ""
	var result = isSubsequence(s, t)
	fmt.Println(result)
}

//
// isSubsequence
// CPU：100%
// 内存：69.92%
//
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}

//
// isSubsequence1
// CPU：100%
// 内存：100%
//
func isSubsequence1(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	var mainLen = len(t)
	var subLen = len(s)
	var mainLeft = 0
	var mainRight = mainLen - 1
	var subLeft = 0
	var subRight = subLen - 1
	var status int8 = 0
	for mainLeft <= mainRight {
		if s[subLeft:subLeft+1] == t[mainLeft:mainLeft+1] {
			mainLeft++
			subLeft++
			status = 1
		}
		if s[subRight:subRight+1] == t[mainRight:mainRight+1] {
			mainRight--
			subRight--
			status = 2
		}
		if status == 0 {
			mainLeft++
			mainRight--
		} else if status == 1 {
			mainRight--
		} else {
			mainLeft++
		}
		if subLeft > subRight {
			return true
		}
	}
	return false
}
