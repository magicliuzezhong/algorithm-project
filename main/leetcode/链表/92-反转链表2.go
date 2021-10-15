//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/15 10:25
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var node = createListNode3()
	var result = reverseBetween(node, 1, 1)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right || head == nil {
		return head
	}
	var second *ListNode = nil
	var p = head
	for i := 0; i < right; i++ {
		if i == right-1 {
			second = p.Next
			p.Next = nil
			break
		}
		p = p.Next
	}
	p = head
	var index = 1
	for index < left-1 {
		p = p.Next
		index++
	}
	if left == 1 {
		var mid, midTail = reversListNode(head)
		midTail.Next = second
		return mid
	} else {
		var mid, midTail = reversListNode(p.Next)
		p.Next = mid
		midTail.Next = second
	}
	return head
}

func reversListNode(head *ListNode) (*ListNode, *ListNode) {
	var cur = head
	var pre *ListNode = nil
	for cur != nil {
		var next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre, head
}

func createListNode3() *ListNode {
	var root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	//root.Next.Next = &ListNode{Val: 3}
	//root.Next.Next.Next = &ListNode{Val: 4}
	//root.Next.Next.Next.Next = &ListNode{Val: 5}
	return root
}
