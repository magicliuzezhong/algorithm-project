//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/9 15:59
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)

	go func() {
		var index = 0
		for {
			time.Sleep(time.Second * 1)
			ch <- index
			if index == 100 {
				close(ch)
				return
			}
			index += 10
		}
	}()

	var timers = time.After(time.Second * 15)
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(i)
		case <-timers:
			fmt.Println("执行超时")
			return
		}
	}

	//var arr = []int{1, 2, 3, 4, 5}
	//fmt.Println(arr)
	//arr = append(arr[:1], arr[2:]...)
	//fmt.Println(arr)
	//fmt.Println(arr[1:3])

}

func timeTest() {
	var arr = createArr()
	var totalTime int64 = 0
	var totalTime1 int64 = 0
	for i := 0; i < 1000; i++ {
		var needTime = testReadArr1(arr)
		totalTime += needTime
	}
	for i := 0; i < 1000; i++ {
		var needTime1 = testReadArr2(arr)
		totalTime1 += needTime1
	}
	fmt.Println("平均时间", totalTime/1000)  //平均时间 1308340
	fmt.Println("平均时间", totalTime1/1000) //平均时间 13268588

	//平均时间 1313864
	//平均时间 1332086
	//平均时间 11005392
	//平均时间 10130807
}

func testReadArr1(arr [][]int) int64 {
	var start = time.Now().UnixNano()
	for i := 0; i < 1024; i++ {
		for j := 0; j < 1024; j++ {
			arr[i][j] = 1
		}
	}
	var end = time.Now().UnixNano()
	return end - start
}
func testReadArr2(arr [][]int) int64 {
	var start = time.Now().UnixNano()
	for i := 0; i < 1024; i++ {
		for j := 0; j < 1024; j++ {
			arr[j][i] = 1
		}
	}
	var end = time.Now().UnixNano()
	return end - start
}

func createArr() [][]int {
	var arr [][]int
	for i := 0; i < 1024; i++ {
		var row = make([]int, 1024)
		arr = append(arr, row)
	}
	return arr
}
