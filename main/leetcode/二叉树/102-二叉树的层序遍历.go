//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/25 14:06
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var tree = GetSymmetricTree()
	var result = levelOrder(tree)
	fmt.Println(result)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result = make([][]int, 0)
	var queue = []*TreeNode{root}
	var queueLen = 0
	for len(queue) > 0 {
		queueLen = len(queue)
		var tmpRes = make([]int, 0)
		for _, val := range queue {
			tmpRes = append(tmpRes, val.Val)
			if val.Left != nil {
				queue = append(queue, val.Left)
			}
			if val.Right != nil {
				queue = append(queue, val.Right)
			}
		}
		result = append(result, tmpRes)
		queue = queue[queueLen:]
	}
	return result
}
