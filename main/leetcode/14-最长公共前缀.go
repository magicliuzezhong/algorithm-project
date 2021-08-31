//
// Package leetcode
// @Description：
// @Author：liuzezhong 2021/7/29 16:16
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var strs = []string{"flower", "flow", "flight"}
	//var strs = []string{"a", "ab"}
	fmt.Println(lcp("", ""))
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	var basic = strs[0]
	for i := 1; i < len(strs); i++ {
		basic = lcp(basic, strs[i])
	}
	return basic
}

func lcp(left, right string) string {
	var preIndex = 1
	for preIndex <= len(left) && preIndex <= len(right) {
		if left[:preIndex] != right[:preIndex] {
			break
		}
		preIndex++
	}
	return left[:preIndex - 1]
}
