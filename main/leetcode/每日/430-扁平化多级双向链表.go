//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/24 10:41
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var root = createMultiNode()
	var result = flatten(root)
	//var result = createMultiNode()
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func createMultiNode() *Node {
	var node1 = &Node{
		Val: 1,
	}
	var node2 = &Node{
		Val: 2,
	}
	var node3 = &Node{
		Val: 3,
	}
	var node4 = &Node{
		Val: 4,
	}
	var node5 = &Node{
		Val: 5,
	}
	var node6 = &Node{
		Val: 6,
	}
	var node7 = &Node{
		Val: 7,
	}
	var node8 = &Node{
		Val: 8,
	}
	var node9 = &Node{
		Val: 9,
	}
	var node10 = &Node{
		Val: 10,
	}
	var node11 = &Node{
		Val: 11,
	}
	var node12 = &Node{
		Val: 12,
	}
	// 处理连接关系
	node1.Next = node2
	node2.Prev = node1
	node2.Next = node3
	node3.Prev = node2
	node3.Next = node4
	node4.Prev = node3
	node4.Next = node5
	node5.Prev = node4
	node5.Next = node6
	node6.Prev = node5

	node7.Next = node8
	node8.Prev = node7
	node8.Next = node9
	node9.Prev = node8
	node9.Next = node10
	node10.Prev = node9

	node11.Next = node12
	node12.Prev = node11

	//node1.Child = node7
	node3.Child = node7
	node8.Child = node11

	return node1
}

//
// flatten
// 第一时间就想到用栈来实现（-_-）就很nice，cpu：100%，内存：100%
//
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	var stack = []*Node{root}
	var rootResult = &Node{}
	var resultPoint = rootResult
	for len(stack) > 0 {
		var top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		resultPoint.Next = top
		top.Prev = resultPoint
		resultPoint = resultPoint.Next
		if top.Next != nil {
			stack = append(stack, top.Next)
		}
		if top.Child != nil {
			stack = append(stack, top.Child)
			top.Child = nil
		}
	}
	rootResult.Next.Prev = nil
	return rootResult.Next
}
