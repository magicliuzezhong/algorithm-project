//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/26 09:57
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var tree = GetSymmetricTree()
	var result = maxDepth(tree)
	fmt.Println(result)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
