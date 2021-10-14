//
// Package 链表
// @Description：
// @Author：liuzezhong 2021/10/14 16:22
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var node = createNode1()
	//var result, _ = reverseListNode(node)
	var result = rotateRight(node, 5)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	// 第一步查找到一共有多少个节点
	if head == nil {
		return head
	}
	var count = 0
	var p = head
	var tail *ListNode = nil
	for p != nil {
		if p.Next == nil {
			tail = p
		}
		p = p.Next
		count++
	}
	k = k % count
	if k == 0 {
		return head
	}
	var fast = head
	var slow = head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		if fast.Next == nil {
			var newHead = slow.Next
			slow.Next = nil // 断开最后的节点
			tail.Next = head
			return newHead
		}
		fast = fast.Next
		slow = slow.Next
	}
	return head
}

func createNode1() *ListNode {
	var root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	root.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	return root
}
