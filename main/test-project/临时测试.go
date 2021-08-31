//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/25 09:41
// @Company cloud-ark.com
//
package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	aaa, err := testErr1()
	fmt.Println(aaa)
	fmt.Println(err.Error())

	if err := testErr(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
}

func testErr() error  {
	return errors.New("aaa")
}

func testErr1() (int, error) {
	return 1, errors.New("bbb")
}
