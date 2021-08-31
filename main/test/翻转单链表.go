//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/9 09:18
// @Company cloud-ark.com
//
package main

import "fmt"

type Linklist struct {
	val  int
	next *Linklist
}

func main() {
	// 1 -> 2 -> 3 -> 4 -> 5
	// 5 -> 4 -> 3 -> 2 -> 1
	var l = getLinklist()
	printLinklist(l)
	//var a = revertLinkList(l)
	//printLinklist(a)
	var a = revertLinkList1(l)
	printLinklist(a)
}

func revertLinkList(l *Linklist) *Linklist {
	var front *Linklist = nil
	var point = l
	for point != nil {
		var next = point.next
		point.next = front
		front = point
		point = next
	}
	return front
}

func revertLinkList1(l *Linklist) *Linklist {
	if l == nil || l.next == nil {
		return l
	}
	var newHead = revertLinkList1(l.next)
	l.next.next = l
	l.next = nil
	return newHead
}

func printLinklist(l *Linklist) {
	for l != nil {
		fmt.Printf("%d  ", l.val)
		l = l.next
	}
	fmt.Println()
}

func getLinklist() *Linklist {
	var l = &Linklist{
		val: 1,
	}
	l.next = &Linklist{val: 2}
	l.next.next = &Linklist{val: 3}
	l.next.next.next = &Linklist{val: 4}
	l.next.next.next.next = &Linklist{val: 5}
	return l
}
