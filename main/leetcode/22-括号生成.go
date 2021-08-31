//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/19 15:35
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var result = generateParenthesis(1)
	fmt.Println(result)
}

//
// generateParenthesis
// 生成有效的括号，最先想到的思路，没思路（-_-）GG
// 之后想了下，决定还是使用回溯，回溯的条件是啥，又没想清楚 =<-_->=
// 看了下题解，牛逼
// 最终解决思路 -> 使用回溯算法，进行数据回溯，
// 回溯的关键点 保证不会出现外括号不能在外面 close < open 保证不会出现 )(，)()(,
// 感想 -> 根本想不出来
//
func generateParenthesis(n int) []string {
	var result = make([]string, 0)
	backtrace(&result, "", 0, 0, n)
	return result
}

func backtrace(result *[]string, data string, open int, close int, n int) {
	if len(data) == 2*n {
		*result = append(*result, data)
		return
	}
	if open < n {
		data += "("
		var dataLen = len(data)
		backtrace(result, data, open+1, close, n)
		data = data[:dataLen-1] + data[dataLen:]
	}
	if close < open {
		data += ")"
		var dataLen = len(data)
		backtrace(result, data, open, close+1, n)
		data = data[:dataLen-1] + data[dataLen:]
	}
}
