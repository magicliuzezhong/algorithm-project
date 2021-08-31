//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/17 18:13
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	//var a = test()
	//fmt.Println(a)
	//test1()
	var t = test2()
	var c = t(1)
	fmt.Println(c)
	var a = make(map[int]string, 10)

	a[1] = "v"
	a[2] = "v"
	fmt.Println(len(a))

	var n = 7  // 00111
	fmt.Println(n)

	// (1 + n) * n / 2 = 1/2n + 1/2 n~2
	// (1 + 100) * 100 / 2
	// 101 * 50 = 5050
}

func test2() func(a int) int {
	var c = 2
	var inner = func(b int) int {
		c++
		return b + c
	}
	return inner
}

func test1() {
	var a = 1
	var b = func() {
		a++
	}
	b()
	b()
	fmt.Println(a)
}

func test() int {
	var a = 1
	defer func() {
		a++
	}()
	return a
}
