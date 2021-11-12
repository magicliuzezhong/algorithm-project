//
// Package main
// @Description：
// @Author：liuzezhong 2021/11/12 17:46
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var p = "hello"
	var t = "lo"
	var index = strStr(p, t)
	fmt.Println(index)
}

func strStr(haystack string, needle string) int {
	var pLen = len(haystack)
	var tLen = len(needle)
	switch {
	case tLen == 0:
		return 0
	case pLen < tLen:
		return -1
	case pLen == tLen:
		if haystack == needle {
			return 0
		}
	case pLen > tLen:
		var pIndex = 0
		var tIndex = 0
		for pIndex < pLen {
			var initIndex = pIndex
			for pIndex < pLen && tIndex < tLen && haystack[pIndex] == needle[tIndex] {
				pIndex++
				tIndex++
			}
			if tIndex == tLen {
				return initIndex
			}
			pIndex = initIndex + 1
			tIndex = 0
		}
	}
	return -1
}
