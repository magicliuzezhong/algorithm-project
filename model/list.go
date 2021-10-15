//
// Package model
// @Description：
// @Author：liuzezhong 2021/8/19 16:53
// @Company cloud-ark.com
//
package model

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}
