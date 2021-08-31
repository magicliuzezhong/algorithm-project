//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/19 11:01
// @Company cloud-ark.com
//
package main

import "fmt"
import "../../model"

func main() {
	var root = &model.ListNode{
		Val: 1,
	}
	root.Next = &model.ListNode{
		Val:  2,
		Next: root,
	}
	var result = detectCycle(root)
	fmt.Println(result)
}

func detectCycle(head *model.ListNode) *model.ListNode {
	if head == nil {
		return nil
	}
	var fastNode = head
	var slowNode = head
	for {
		if fastNode.Next == nil || fastNode.Next.Next == nil {
			return nil
		}
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
		if fastNode == slowNode {
			break
		}
	}
	slowNode = head
	for slowNode != fastNode {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}
	return slowNode
}
