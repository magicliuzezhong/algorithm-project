//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/24 18:26
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var node = createTree3()

	var a = preorderTraversal(node)
	fmt.Println(a)
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result = make([]int, 0)
	var stack = []*TreeNode{root}
	for len(stack) > 0 {
		var top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return result
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result = make([]int, 0)
	result = append(result, root.Val)
	var leftResult = preorderTraversal(root.Left)
	var rightResult = preorderTraversal(root.Right)
	if len(leftResult) > 0 {
		result = append(result, leftResult...)
	}
	if len(rightResult) > 0 {
		result = append(result, rightResult...)
	}
	return result
}

func createTree3() *TreeNode {
	var root = &TreeNode{
		Val: 5,
	}
	root.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 4,
	}
	root.Right.Left = &TreeNode{
		Val: 3,
	}
	root.Right.Right = &TreeNode{
		Val: 6,
	}
	return root
}
