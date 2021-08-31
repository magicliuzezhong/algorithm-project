//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/24 11:42
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var result = generateTrees(3)
	fmt.Println(len(result))
	for _, node := range result {
		PrintTree(node)
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var result = make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		var leftTree = helper(start, i-1)
		var rightTree = helper(i+1, end)
		for _, left := range leftTree {
			for _, right := range rightTree {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return result
}
