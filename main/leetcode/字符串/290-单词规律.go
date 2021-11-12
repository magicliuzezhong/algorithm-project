//
// Package main
// @Description：
// @Author：liuzezhong 2021/11/12
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"strings"
)

func main() {
	//var pattern = "abba"
	//var str = "dog cat cat dog"
	var pattern = "aaaa"
	var str = "dog cat cat dog"
	var res = wordPattern(pattern, str)
	fmt.Println(res)
}

func wordPattern(pattern string, s string) bool {
	var strs = strings.Split(s, " ")
	var patternLen = len(pattern)
	if patternLen != len(strs) {
		return false
	}
	var record1 = make(map[uint8]int)
	var record2 = make(map[string]int)
	for i := 0; i < patternLen; i++ {
		if record1[pattern[i]] != 0 && strs[record1[pattern[i]]-1] != strs[i] {
			return false
		}
		if record2[strs[i]] != 0 && pattern[i] != pattern[record2[strs[i]]-1] {
			return false
		}
		record1[pattern[i]] = i + 1
		record2[strs[i]] = i + 1
	}
	return true
}
