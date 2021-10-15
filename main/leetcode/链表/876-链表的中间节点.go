//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/15 11:17
// @Company cloud-ark.com
//
package main

import . "../../../model"

func main() {
	var node = createListNode5()
	var result = middleNode(node)
	PrintListNode(result)
}

func middleNode(head *ListNode) *ListNode {
	var fast = head
	var slow = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func createListNode5() *ListNode {
	var root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	return root
}
