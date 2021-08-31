//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/26 10:04
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
	"math"
)

func main() {
	//var tree = GetSymmetricTree()
	var tree = CreateTree1()

	var result = isBalanced(tree)
	fmt.Println(result)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return int(math.Abs(float64(getTreeHeight(root.Left)-getTreeHeight(root.Right)))) <= 1 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right)
}

func getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(getTreeHeight(root.Left)), float64(getTreeHeight(root.Right)))) + 1
}
