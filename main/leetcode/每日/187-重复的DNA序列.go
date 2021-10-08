//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/8 15:33
// @Company cloud-ark.com
//
package main

import (
	"fmt"
)

func main() {
	//var s = "AAAAAAAAAAAAA"
	//var s = "AAAAAAAAAAA"
	var s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	var result = findRepeatedDnaSequences(s)
	fmt.Println(result)
}

func findRepeatedDnaSequences(s string) []string {
	var record = make(map[string]int)
	var sLen = len(s)
	var result = make([]string, 0)
	for i := 0; i <= sLen-10; i++ {
		var word = s[i : i+10]
		if val, ok := record[word]; ok {
			val++
			record[word] = val
			if val == 2 {
				result = append(result, word)
			}
		} else {
			record[word] = 1
		}
	}
	return result
}
