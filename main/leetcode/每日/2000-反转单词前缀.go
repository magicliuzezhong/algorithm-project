// Package main
// @Author 刘泽忠
// @Time 2022-02-05 21:03:33
// @Description: 反转单词前缀
package main

import (
	"fmt"
)

func main() {
	prefix := reversePrefix("abcdefd", 'd')
	fmt.Println(prefix)
}

func reversePrefix(word string, ch byte) string {
	for i, val := range word {
		if byte(val) == ch {
			return reverse(word[:i+1]) + word[i+1:]
		}
	}
	return word
}

func reverse(word string) string {
	res := make([]int32, len(word))
	start, end := 0, len(word)-1
	for i, val := range word {
		res[i] = val
	}
	for start < end {
		res[start], res[end] = res[end], res[start]
		start++
		end--
	}
	return string(res)
}
