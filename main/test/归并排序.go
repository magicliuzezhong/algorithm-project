//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/6 11:12
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{9, 8, 7, 6, 5}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 9, 8, 7, 6, 5
// 9, 8,    7, 6, 5
// 9,   8,    7,   6, 5
// 9,   8,    7,   6,  5
// 8, 9,    7,   5,  6
// 8, 9,    5,  6, 7
// 5,  6, 7, 8, 9

func mergeData(arr []int, left, mid, right int) {
	//7, 6, 5    left = 0 mid = 1 right = 2
	var arrLen = right - left + 1
	var newArr = make([]int, arrLen)
	var basicLeft = left
	var basicMid = mid + 1
	var index = 0
	for basicLeft <= mid && basicMid <= right {
		if arr[basicLeft] < arr[basicMid] {
			newArr[index] = arr[basicLeft]
			basicLeft++
		} else {
			newArr[index] = arr[basicMid]
			basicMid++
		}
		index++
	}
	for basicLeft <= mid {
		newArr[index] = arr[basicLeft]
		basicLeft++
		index++
	}
	for basicMid <= right {
		newArr[index] = arr[basicMid]
		basicMid++
		index++
	}
	for i := 0; i < arrLen; i++ {
		arr[left + i] = newArr[i]
	}
}

func mergeSort(arr []int, left, right int) {
	if left < right {
		var mid = (left + right) / 2
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		mergeData(arr, left, mid, right)
	}
}
