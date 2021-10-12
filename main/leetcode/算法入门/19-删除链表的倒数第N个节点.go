//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/12 16:10
// @Company cloud-ark.com
//
package main

import . "../../../model"
import "fmt"

func main() {
	var node = createListNode1()
	var res = removeNthFromEnd(node, 2)
	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 移动n个节点
	var slow = head
	var fast = head
	for i := 0; i < n; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	if fast == nil {
		head = head.Next
	}
	for fast != nil {
		if fast.Next == nil {
			slow.Next = slow.Next.Next
		}
		fast = fast.Next
		slow = slow.Next
	}
	return head
}

func createListNode1() *ListNode {
	var root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	return root
}
