//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/9 16:19
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var s = "Let's take LeetCode contest "
	//var s = " a"
	fmt.Println(s)
	var result = reverseWords(s)
	fmt.Println(result)
}

func reverseWords(s string) string {
	var sLen = len(s)
	var result = make([]byte, sLen)
	var resIndex = 0
	var index = 0
	for index < sLen {
		if s[index] == ' ' {
			result[resIndex] = ' '
			resIndex++
			index++
			continue
		}
		var end = index
		for j := index + 1; j < sLen; j++ {
			if s[j] != ' ' {
				end = j
			} else {
				break
			}
		}
		for j := end; j >= index; j-- {
			result[resIndex] = s[j]
			resIndex++
		}
		index = end + 1
	}
	return string(result)
}
