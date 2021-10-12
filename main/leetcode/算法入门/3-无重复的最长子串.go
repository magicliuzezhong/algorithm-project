//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/12 16:27
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var s = "abcabcbb"
	var s = "pwwkew"
	var res = lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	// a b c a b c b b
	// 1 1 1 1
	var sLen = len(s)
	var left = 0
	var right = 0
	var record = make(map[byte]int)
	var res = 0
	for right < sLen {
		var char = s[right]
		if val, ok := record[char]; ok {
			if val >= left {
				left = val + 1
			}
		}
		record[char] = right
		if right-left+1 > res {
			res = right - left + 1
		}
		right++
	}
	return res
}
