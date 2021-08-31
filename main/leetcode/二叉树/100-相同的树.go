//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/25 10:30
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var tree = CreateTree1()
	var tree1 = CreateTree1()
	var result = isSameTree(tree, tree1)
	fmt.Println(result)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
