//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/15 18:15
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	// 1 2 3 4 5 6 8 9 10
	//     1
	//   2
	//       1
	//     3
	var result = numTrees(6)
	fmt.Println(result)
}

func numTrees(n int) int {
	if n == 2 || n == 1 {
		return n
	}
	// 1 2 3 4 5 6
	// 3
	var val = 0
	for i := 1; i <= n; i++ {
		var leftCount = numTrees(i - 1)
		var rihitCount = numTrees(n - i)
		if leftCount == 0 {
			val += rihitCount
		}
		if rihitCount == 0 {
			val += leftCount
		}
		val += leftCount * rihitCount
	}
	return val
}
