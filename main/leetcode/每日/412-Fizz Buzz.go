//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/13 15:13
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var target = 15
	var res = fizzBuzz(target)
	fmt.Println(res)
}

func fizzBuzz(n int) []string {
	var res = make([]string, 0)
	for i := 1; i <= n; i++ {
		var a = i%3 == 0
		var b = i%5 == 0
		if a && b {
			res = append(res, "FizzBuzz")
		} else if a {
			res = append(res, "Fizz")
		} else if b {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
