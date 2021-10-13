//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/13 18:10
// @Company cloud-ark.com
//
package main

import . "../../../model"

func main() {
	var node1 = createNode1()
	var node2 = createNode2()
	var result = mergeTrees(node1, node2)
	PrintTree(result)
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	var left = mergeTrees(root1.Left, root2.Left)
	var right = mergeTrees(root1.Right, root2.Right)
	root1.Left = left
	root1.Right = right
	return root1
}

func createNode1() *TreeNode {
	var root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	return root
}
func createNode2() *TreeNode {
	var root = &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}
