package main

import . "../../../model"

func main() {
	var head = createList()
	var delNode = head.Next.Next
	deleteNode(delNode)
	PrintListNode(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	*node = *node.Next
}

func createList() *ListNode {
	var root = &ListNode{Val: 4}
	root.Next = &ListNode{Val: 5}
	root.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next = &ListNode{Val: 9}
	return root
}
