//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/12 16:01
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
)

func main() {
	var node = createListNode()
	var result = middleNode(node)
	fmt.Println(result)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var fast = head
	var slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func createListNode() *ListNode {
	var root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	//root.Next.Next.Next = &ListNode{Val: 4}
	//root.Next.Next.Next.Next = &ListNode{Val: 5}
	//root.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	return root
}
