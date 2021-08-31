//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/23 11:44
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{1,2,3}
	testArr1(arr)
	testArr2(&arr)
	fmt.Println("  ==== 最后结果 ====")
	fmt.Println(arr, ", len：", len(arr), ", cap: ", cap(arr))
}

func testArr1(arr []int)  {
	fmt.Println(" ===== 测试非指针类型 ===== ")
	fmt.Println(arr, ", len：", len(arr), ", cap: ", cap(arr))
	arr[0] = -1
	arr[1] = -9
	arr = append(arr, 5)
	fmt.Println(arr, ", len：", len(arr), ", cap: ", cap(arr))
}

func testArr2(arr *[]int)  {
	fmt.Println(" ===== 测试指针类型 ===== ")
	fmt.Println(arr, ", len：", len(*arr), ", cap: ", cap(*arr))
	(*arr)[0] = -2
	*arr = append(*arr, 4)
	*arr = append(*arr, 5)
	fmt.Println(arr, ", len：", len(*arr), ", cap: ", cap(*arr))
}
