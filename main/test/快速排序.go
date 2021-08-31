//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/6 10:26
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{4, 1, 5, 1, 9, 3, 4, 7}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	var basic = arr[left]
	var leftIndex = left
	var rightIndex = right
	for leftIndex < rightIndex {
		for leftIndex < rightIndex && arr[rightIndex] > basic {
			rightIndex--
		}
		for leftIndex < rightIndex && arr[leftIndex] <= basic {
			leftIndex++
		}
		if leftIndex <= rightIndex {
			arr[leftIndex], arr[rightIndex] = arr[rightIndex], arr[leftIndex]
		}
	}
	arr[left] = arr[rightIndex]
	arr[rightIndex] = basic
	quickSort(arr, left, leftIndex-1)
	quickSort(arr, rightIndex+1, right)
}
