//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/15 11:42
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var result = countAndSay(29)
	fmt.Println(result)
}

// 1
// 11
// 21
// 1211
// 111221
// 312211
// 13112221
// 1113213211
// 31131211131221
// 13211311123113112111

// @Description: 后面一个值只跟上一个值有关，所以可以使用滚动列表优化
func countAndSay(n int) string {
	var oldStr = "1"
	var newStr = strings.Builder{}
	for i := 1; i < n; i++ {
		var oldLen = len(oldStr)
		var index = 0
		newStr.Reset()
		for index < oldLen {
			var basic = oldStr[index : index+1]
			var count = 1
			for j := index + 1; j < oldLen; j++ {
				if oldStr[j:j+1] == basic {
					index++
					count++
				} else {
					break
				}
			}
			newStr.WriteString(strconv.Itoa(count))
			newStr.WriteString(basic)
			index++
		}
		oldStr = newStr.String()
	}
	return oldStr
}

//
// countAndSay
// @Description: 后面一个值只跟上一个值有关，所以可以使用滚动列表优化
//
func countAndSay1(n int) string {
	if n == 1 {
		return "1"
	}
	var res = countAndSay(n - 1)
	var resLen = len(res)
	var result = ""
	var index = 0
	for index < resLen {
		var basic = res[index : index+1]
		var count = 1
		for i := index + 1; i < resLen; i++ {
			if res[i:i+1] == basic {
				index++
				count++
			} else {
				break
			}
		}
		result += strconv.Itoa(count)
		result += basic
		index++
	}
	return result
}
