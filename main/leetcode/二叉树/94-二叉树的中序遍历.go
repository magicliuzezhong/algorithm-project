//
// Package 二叉树
// @Description：
// @Author：liuzezhong 2021/8/24 10:58
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var result = inorderTraversal(CreateTree1())
	fmt.Println(result)
}

func inorderTraversal(root *TreeNode) []int {
	var stack = []*TreeNode{}
	var res = make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	var result = make([]int, 0)
	centerTraversal(root, &result)
	return result
}

func centerTraversal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	centerTraversal(root.Left, result)
	*result = append(*result, root.Val)
	centerTraversal(root.Right, result)
}
