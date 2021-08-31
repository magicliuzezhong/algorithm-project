//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/19 10:20
// @Company cloud-ark.com
//
package main

import "fmt"

type cirLink struct {
	val  int
	next *cirLink
}

func main() {
	var link = createLink()
	var isC = isCir(link)
	fmt.Println(isC)
	var a = getEntry(link)
	fmt.Println(a)
}

func getEntry(link *cirLink) *cirLink {
	var fastLink = link
	var slowLink = link
	for {
		if fastLink.next == nil || fastLink.next.next == nil {
			return nil
		}
		fastLink = fastLink.next.next
		slowLink = slowLink.next
		if fastLink == slowLink {
			 break
		}
	}
	slowLink = link
	for {
		if fastLink == slowLink {
			return fastLink
		}
		slowLink = slowLink.next
		fastLink = fastLink.next
	}
}

// 快慢指针
func isCir(link *cirLink) bool {
	if link == nil {
		return false
	}
	var slowLink = link
	var fastLink = link
	for {
		if fastLink.next == nil || fastLink.next.next == nil {
			return false
		}
		slowLink = slowLink.next
		fastLink = fastLink.next.next
		if fastLink == slowLink {
			return true
		}
	}
}

func createLink() *cirLink {
	var root = &cirLink{
		val: 1,
	}
	root.next = &cirLink{
		val: 2,
	}
	root.next.next = &cirLink{
		val: 3,
	}
	root.next.next.next = &cirLink{
		val: 4,
	}
	root.next.next.next.next = &cirLink{
		val: 5,
	}
	root.next.next.next.next.next = root.next
	return root
}
