//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/17 16:11
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var result = letterCombinations("")
	fmt.Println(len(result))
	fmt.Println(result)
}

//
// letterCombinations
// 比较简单的回溯
//
func letterCombinations(digits string) []string {
	var result = make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	var keys = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	backtraceWords(digits, keys, 0, len(digits), "", &result)
	return result
}

func backtraceWords(digits string, keys map[string]string, start int, wordLen int, data string, result *[]string) {
	if len(data) == wordLen {
		*result = append(*result, data)
	}
	for i := start; i < wordLen; i++ {
		var word = digits[i : i+1]
		if wordTable, ok := keys[word]; ok {
			for j := 0; j < len(wordTable); j++ {
				backtraceWords(digits, keys, i+1, wordLen, data+wordTable[j:j+1], result)
			}
		}
	}
}
