//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/19 16:52
// @Company cloud-ark.com
//
package main

import (
	"../../model"
	"fmt"
)

func main() {
	var listNode = createNode()
	//var result = removeNthFromEnd(listNode, 1)
	var result = removeNthFromEnd1(listNode, 1)
	fmt.Println(result)
	query(result)
}

func removeNthFromEnd1(listNode *model.ListNode, n int) *model.ListNode {
	var shell = backtrace19(listNode, n)
	if shell == n {
		return listNode.Next
	}
	return listNode
}

func backtrace19(listNode *model.ListNode, n int) int {
	if listNode == nil {
		return 0
	}
	var shell = backtrace19(listNode.Next, n)
	if shell == n {
		listNode.Next = listNode.Next.Next
	}
	return shell + 1
}

func query(listNode *model.ListNode) {
	var p = listNode
	for p != nil {
		fmt.Print(" ", p.Val)
		p = p.Next
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
// removeNthFromEnd
// 快慢指针，非常经典的问题
// 不用看答案我都做出来了=v=
//
func removeNthFromEnd(head *model.ListNode, n int) *model.ListNode {
	var firstNode = head
	var endNode = head
	for i := 0; i < n; i++ {
		if endNode == nil {
			return nil
		}
		endNode = endNode.Next
	}
	if endNode == nil {
		head = head.Next
	}
	for endNode != nil {
		if endNode.Next == nil {
			firstNode.Next = firstNode.Next.Next
			break
		}
		firstNode = firstNode.Next
		endNode = endNode.Next
	}
	return head
}

func createNode() *model.ListNode {
	var root = &model.ListNode{
		Val: 1,
	}
	//root.Next = &model.ListNode{
	//	Val: 2,
	//}
	//root.Next.Next = &model.ListNode{
	//	Val: 3,
	//}
	//root.Next.Next.Next = &model.ListNode{
	//	Val: 4,
	//}
	//root.Next.Next.Next.Next = &model.ListNode{
	//	Val: 5,
	//}
	//root.Next.Next.Next.Next.Next = &model.ListNode{
	//	Val: 6,
	//}
	//root.Next.Next.Next.Next.Next.Next = &model.ListNode{
	//	Val: 7,
	//}
	return root
}
