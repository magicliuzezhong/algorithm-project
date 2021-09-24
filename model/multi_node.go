//
// Package model
// @Description：
// @Author：liuzezhong 2021/9/24 10:39
// @Company cloud-ark.com
//
package model

//
// Node
// @Description: 多级链表
//
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}
