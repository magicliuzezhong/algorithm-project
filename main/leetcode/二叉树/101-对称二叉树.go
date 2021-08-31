//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/25 10:59
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var tree = GetSymmetricTree()
	var result = isSymmetric(tree)
	fmt.Println(result)
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricNode(root, root)
}

func isSymmetricNode(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isSymmetricNode(left.Left, right.Right) && isSymmetricNode(left.Right, right.Left)
}
