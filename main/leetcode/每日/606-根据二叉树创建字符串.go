package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	res := tree2str(root)
	fmt.Println(res)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	switch {
	case root == nil:
		return ""
	case root.Left == nil && root.Right == nil:
		return fmt.Sprintf("%d", root.Val)
	case root.Right == nil: // 右子树为空
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	default:
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	}
}
