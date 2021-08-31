//
// Package model
// @Description：
// @Author：liuzezhong 2021/8/24 10:58
// @Company cloud-ark.com
//
package model

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree1() *TreeNode {
	var root = &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Left.Left = &TreeNode{
		Val: 4,
	}
	root.Left.Right = &TreeNode{
		Val: 5,
	}
	root.Left.Left.Left = &TreeNode{
		Val: 6,
	}
	return root
}

func GetMinDepthTree() *TreeNode {
	var root = &TreeNode{
		Val: 3,
	}
	root.Left = &TreeNode{
		Val: 9,
	}
	root.Right = &TreeNode{
		Val: 20,
	}
	root.Right.Left = &TreeNode{
		Val: 15,
	}
	root.Right.Right = &TreeNode{
		Val: 15,
	}
	return root
}

func GetMinDepthTree1() *TreeNode {
	var root = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Right.Right = &TreeNode{
		Val: 4,
	}
	root.Right.Right.Right = &TreeNode{
		Val: 5,
	}
	root.Right.Right.Right.Right = &TreeNode{
		Val: 6,
	}
	return root
}

// GetSymmetricTree 获取对称二叉树
func GetSymmetricTree() *TreeNode {
	var root = &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 2,
	}
	root.Left.Left = &TreeNode{
		Val: 3,
	}
	root.Left.Right = &TreeNode{
		Val: 4,
	}
	root.Right.Left = &TreeNode{
		Val: 4,
	}
	root.Right.Right = &TreeNode{
		Val: 3,
	}
	return root
}

//
// InorderTraversal
// 二叉树中序遍历
//
func InorderTraversal(root *TreeNode) []*TreeNode {
	var result = make([]*TreeNode, 0)
	var stack = make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root)
		root = root.Right
	}
	return result
}

//
// PrintTree
// 二叉树打印（只为美观）
// 该代码抄自C++版本，(-_-)感觉我已经入门C++了呀
// 参考地址：https://www.tqwba.com/x_d/jishu/288287.html
//
func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	var tmpNode = root
	var intv = InorderTraversal(tmpNode)
	var tmp_str = ""
	var location = 0
	var recordFirst = make(map[*TreeNode]int)
	for _, val := range intv {
		location = len(tmp_str)
		tmp_str += strconv.Itoa(val.Val) + " "
		recordFirst[val] = location
	}
	var tmpstrLen = len(tmp_str)
	for i := 0; i < tmpstrLen; i++ { //全部置成 " "
		tmp_str = tmp_str[:i] + " " + tmp_str[i+1:]
	}
	var queue = make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		var currentSize = len(queue)
		var cur_loc = 0
		var tmpStr1 = tmp_str
		var tmpStr2 = tmp_str
		for i := 1; i <= currentSize; i++ {
			var node = queue[0]
			queue = append(queue[:0], queue[1:]...)
			cur_loc = recordFirst[node]
			var num_str = strconv.Itoa(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
				var first_loc = recordFirst[node.Left] + 1
				tmpStr2 = tmpStr2[:first_loc] + "/" + tmpStr2[first_loc+1:]
				first_loc++
				for first_loc < cur_loc {
					tmpStr1 = tmpStr1[:first_loc] + "_" + tmpStr1[first_loc+1:]
					first_loc++
				}
			}
			for j := 0; j < len(num_str); j++ {
				tmpStr1 = tmpStr1[:cur_loc] + num_str[j:j+1] + tmpStr1[cur_loc+1:]
				cur_loc++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				var last_loc = recordFirst[node.Right] - 1
				tmpStr2 = tmpStr2[:last_loc] + "\\" + tmpStr2[last_loc+1:]
				for cur_loc < last_loc {
					tmpStr1 = tmpStr1[:cur_loc] + "_" + tmpStr1[cur_loc+1:]
					cur_loc++
				}
			}
		}
		fmt.Println(tmpStr1)
		fmt.Println(tmpStr2)
	}
}
