//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/9 09:54
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func main() {
	var t = CreateTree()
	fmt.Println(TreeHeight(t))
	var isBalances = isBalance(t)
	fmt.Println(isBalances)
}

func isBalance(tree *Tree) bool {
	if tree == nil {
		return true
	}
	return isBalance(tree.Left) && isBalance(tree.Right) && int(math.Abs(float64(TreeHeight(tree.Left)-TreeHeight(tree.Right)))) <= 1
}

func TreeHeight(tree *Tree) int {
	if tree == nil {
		return 0
	}
	return int(math.Max(float64(TreeHeight(tree.Left)), float64(TreeHeight(tree.Right)))) + 1
}

func CreateTree() *Tree {
	var root = &Tree{Val: 0}
	root.Left = &Tree{Val: 1}
	root.Right = &Tree{Val: 2}
	root.Left.Left = &Tree{Val: 3}
	root.Left.Left.Left = &Tree{Val: 7}
	//root.left.left.left.left = &Tree{val: 8}
	root.Left.Right = &Tree{Val: 4}
	root.Right.Left = &Tree{Val: 5}
	root.Right.Right = &Tree{Val: 6}
	return root
}
