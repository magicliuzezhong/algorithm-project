//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/9 16:09
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var s = []byte{
		'h', 'e', 'l', 'l', 'o',
	}
	reverseString(s)
	fmt.Println(fmt.Sprintf("%s", s))

}

func reverseString(s []byte) {
	var left, right = 0, len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
