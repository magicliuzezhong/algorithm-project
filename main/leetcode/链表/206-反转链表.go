//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/15 09:45
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var node = createListNode2()
	var result = reverseList(node)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	// 1 2 3 4 5 nil
	if head == nil || head.Next == nil {
		return head
	}
	var node = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func reverseList1(head *ListNode) *ListNode {
	var cur = head
	var pre *ListNode = nil
	for cur != nil {
		var next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func createListNode2() *ListNode {
	var root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	return root
}
