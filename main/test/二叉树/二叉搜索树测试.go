//
// Package 二叉树
// @Description：
// @Author：liuzezhong 2021/8/30 10:31
// @Company cloud-ark.com
//
package main

import . "../../../model"

func main() {
	var tree = NewBSTree()
	var arr = []int{8, 3, 5, 4, 1, 2, 6, 13, 7, 10, 9, 12, 15, 14, 16}
	for _, val := range arr {
		tree.Add(val)
	}
	PrintTree(tree.Root)
	tree.Remove(5)
	PrintTree(tree.Root)

	//var c = 2 ^ 3 // 0010 ^ 0011 = 0001
	//var d = 2 ^ 7 // 0010 ^ 0111 = 0101
	//var e = 3 << 1 // 0011 << 1 = 0110 = 6
	//var f = 3 << 2 // 0011 << 2 = 1100 = 12
	//fmt.Println(c) // 0001  =  1
	//fmt.Println(d) // 0101  =  5
	//fmt.Println(e) // 0101  =  6
	//fmt.Println(f) // 0101  =  12
}
