//
// Package 每日
// @Description：
// @Author：liuzezhong 2021/9/18 11:22
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var target = 4
	var result = canWinNim(target)
	fmt.Println(result)
}

//
// canWinNim
// 只要最后给对手剩下4的倍数我稳赢
// 相反，如果当前就是4的倍数那么必输
//
func canWinNim(n int) bool {
	return n%4 != 0
}
