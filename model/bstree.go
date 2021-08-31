//
// Package model
// @Description：二叉搜索树
// @Author：liuzezhong 2021/8/30 10:09
// @Company cloud-ark.com
//
package model

//
// BSTree
// @Description: 二叉搜索树
//
type BSTree struct {
	Root *TreeNode
}

//
// NewBSTree
// @Description: new BST
// @return *BinarySearchTree 二叉搜索树
//
func NewBSTree() *BSTree {
	return &BSTree{}
}

func (b *BSTree) Find(val int) *TreeNode {
	var root = b.Root
	for root != nil {
		if root.Val == val {
			return root
		}
		if root != nil && root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

func (b *BSTree) Add(val int) {
	var newNode = &TreeNode{Val: val}
	if b.Root == nil {
		b.Root = newNode
	} else {
		b.addNode(b.Root, newNode)
	}
}

func (b *BSTree) Remove(val int) bool {
	var node = b.Root
	if node == nil {
		return false
	}
	if node.Val == val {
		node = nil
		return true
	}
	var parent *TreeNode
	var cur = node
	for cur != nil {
		if cur.Val > val {
			parent = cur
			cur = cur.Left
		} else if cur.Val < val {
			parent = cur
			cur = cur.Right
		} else {
			if cur.Left == nil {
				if cur == parent.Left {
					parent.Left = cur.Right
				} else {
					parent.Right = cur.Right
				}
			} else if cur.Right == nil {
				if cur == parent.Left {
					parent.Left = cur.Left
				} else {
					parent.Right = cur.Right
				}
			} else {
				var RightMinLeft = cur.Right
				for RightMinLeft.Left != nil {
					parent = RightMinLeft
					RightMinLeft = RightMinLeft.Left
				}
				cur.Val = RightMinLeft.Val
				if RightMinLeft == parent.Left {
					parent.Left = RightMinLeft.Right
				} else {
					parent.Right = RightMinLeft.Right
				}
				RightMinLeft = nil
			}
			return true
		}
	}
	return false
}

func (b *BSTree) FrontIter() []int {
	var result = make([]int, 0)
	if b.Root == nil {
		return result
	}
	var stack = []*TreeNode{b.Root}
	for len(stack) > 0 {
		var top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return result
}

func (b *BSTree) MidIter() []int {
	var result = make([]int, 0)
	var stack = make([]*TreeNode, 0)
	var root = b.Root
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}

func (b *BSTree) BackIter() []int {
	var result = make([]int, 0)
	if b.Root == nil {
		return result
	}
	var stack = []*TreeNode{b.Root}
	for len(stack) > 0 {
		var top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, top.Val)
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	return result
}

func (b *BSTree) addNode(root, newNode *TreeNode) {
	if root.Val > newNode.Val {
		if root.Left == nil {
			root.Left = newNode
		} else {
			b.addNode(root.Left, newNode)
		}
	} else if root.Val < newNode.Val {
		if root.Right == nil {
			root.Right = newNode
		} else {
			b.addNode(root.Right, newNode)
		}
	}
}
