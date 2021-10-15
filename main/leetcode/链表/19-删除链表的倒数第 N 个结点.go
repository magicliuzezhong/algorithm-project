//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/15 11:00
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
)

func main() {
	var node = createListNode4()
	var result = removeNthFromEnd(node, 1)
	PrintListNode(result)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	var fast = head
	var slow = head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast != nil {
		if fast.Next == nil {
			slow.Next = slow.Next.Next
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	return head
}

func createListNode4() *ListNode {
	var root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	//root.Next.Next = &ListNode{Val: 3}
	//root.Next.Next.Next = &ListNode{Val: 4}
	//root.Next.Next.Next.Next = &ListNode{Val: 5}
	return root
}
