//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/15 11:29
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var node = createListNode6()
	var result = detectCycle(node)
	fmt.Println(result)
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var slow = head
	var fast = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func createListNode6() *ListNode {
	var root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next = root.Next.Next
	return root
}
