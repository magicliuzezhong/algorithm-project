//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/6 10:51
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{9,8,7,6,5,4,3,2,1}
	Heapsort(arr)
	fmt.Println(arr)
}

func Heapsort(arr []int) {
	var lens = len(arr)
	//构建大顶堆
	for i := lens / 2; i >= 0; i-- {
		buildMaxHeap(arr, i, len(arr))
	}
	for i := lens - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		buildMaxHeap(arr, 0, i)
	}
}

func buildMaxHeap(arr []int, parent, len int) {
	for parent*2 < len {
		var maxNode = parent
		if arr[parent*2] > arr[maxNode] {
			maxNode = parent * 2
		}
		if parent*2+1 < len && arr[parent*2+1] > arr[maxNode] {
			maxNode = parent*2 + 1
		}
		if maxNode == parent { //节点相同 退出即可
			break
		}
		arr[parent], arr[maxNode] = arr[maxNode], arr[parent]
		parent = maxNode
	}
}
