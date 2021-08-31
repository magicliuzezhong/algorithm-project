//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/24 17:20
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
	"math"
)

func main() {
	//      4
	//    2   6
	//  1  3 5  7
	//var node = createTree1()
	var node = createTree2()
	PrintTree(node)
	var a = isValidBST(node)
	fmt.Println(a)
}

//
// isValidBST
// 中序遍历
// 利用二叉搜索树进行中序遍历是顺序的这一个特点
// 意味着如果 存在 n > n+1 说明就不是二叉搜索树
//
func isValidBST(root *TreeNode) bool {
	var pre = math.MinInt64
	var stack = make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= pre {
			return false
		}
		pre = root.Val
		root = root.Right
	}
	return true
}

//
// isValidBST1
// 常规递归方式
//
func isValidBST1(root *TreeNode) bool {
	return checkTree(root, math.MinInt64, math.MaxInt64)
}

func checkTree(root *TreeNode, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}
	return checkTree(root.Left, minVal, root.Val) && checkTree(root.Right, root.Val, maxVal)
}

func createTree1() *TreeNode {
	var root = &TreeNode{
		Val: 2,
	}
	root.Left = &TreeNode{
		Val: 1,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	return root
}

func createTree2() *TreeNode {
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
