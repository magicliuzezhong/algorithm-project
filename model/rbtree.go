//
// Package model
// @Description：
// @Author：liuzezhong 2021/8/30 15:42
// @Company cloud-ark.com
//
package model

import (
	"fmt"
	"strconv"
)

type nodeColor bool

const (
	Red   nodeColor = true
	Black nodeColor = false
)

type rbNode struct {
	key    int
	value  interface{}
	parent *rbNode
	left   *rbNode
	right  *rbNode
	color  nodeColor
}

//
// InorderTraversal
// @Description: 中序遍历
//
func (r *rbNode) InorderTraversal() []*rbNode {
	var result = make([]*rbNode, 0)
	var stack = make([]*rbNode, 0)
	var node = r
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node)
		node = node.right
	}
	return result
}

type RBTree struct {
	Root *rbNode
}

func (r *RBTree) getParent(node *rbNode) *rbNode {
	if node != nil {
		return node.parent
	}
	return nil
}

func (r *RBTree) isRed(node *rbNode) bool {
	if node != nil {
		return node.color == Red
	}
	return false
}

func (r *RBTree) isBlack(node *rbNode) bool {
	if node.color == Black {
		return true
	}
	return false
}

func (r *RBTree) setRed(node *rbNode) {
	if node != nil {
		node.color = Red
	}
}

func (r *RBTree) setBlack(node *rbNode) {
	if node != nil {
		node.color = Black
	}
}

func (r *RBTree) Insert(key int, value interface{}) {
	var newNode = &rbNode{
		key:   key,
		value: value,
		color: Red,
	}
	r.insertNode(newNode)
}

func (r *RBTree) insertNode(node *rbNode) {
	var parent *rbNode = nil
	var root = r.Root
	for root != nil {
		parent = root
		if root.key == node.key {
			root.value = node.value
			return
		} else if root.key > node.key {
			root = root.left
		} else {
			root = root.right
		}
	}
	node.parent = parent
	if parent != nil {
		if node.key > parent.key {
			parent.right = node
		} else {
			parent.left = node
		}
	} else {
		r.setBlack(node)
		r.Root = node
	}
	// 插入完成，现在需要处理红黑树平衡性处理
	r.reBalance(node)
	fmt.Println()
}

//
// reBalance
// 插入后需要修复红黑树平衡性问题
// 1、红黑树为空树 -》 将当前节点染色为黑色即可
// 2、插入的key已经存在，无需处理，替换value就行
// 3、插入的节点的父节点为黑色，黑色节点无需变化，依旧是处于平衡状态的，不用调整
//
// 需要处理的是以下情况
// 1、插入的节点的父亲节点是红色
// 1.1、叔叔节点存在，并且为红色（父叔双红）将爸爸和叔叔节点染为黑色，并且再以爷爷为当前节点就行下一轮处理
//
// 1.2、叔叔节点不存在 | 或者为黑色，父亲节点为爷爷节点的左子树
// 1.2.1、插入节点为父亲节点的左子节点（LL）将父亲染为黑色，将爷爷染为红色，然后以爷爷节点右旋，完成
// 1.2.2、插入节点为父亲节点的右子节点（LR）以父亲节点进行一次左旋，得到（LL）然后按照1.2.1处理
//
// 1.3、叔叔节点不存在 | 或者为黑色，父亲节点为爷爷节点的右子树
// 1.3.1、插入的节点为父亲节点的右子树（RR），将父亲染为黑色，爷爷染为红色，最后以爷爷为节点进行左旋，完成
// 1.3.2、插入的节点为父亲节点的左子树（LR），以父亲节点进行一次右旋，得到（RR），然后按照1.3.1处理
//
func (r *RBTree) reBalance(node *rbNode) {
	r.Root.color = Black
	var parent = r.getParent(node)
	var gParent = r.getParent(parent)
	// 需要处理的情景 即父节点是红色
	if parent != nil && r.isRed(parent) {
		// 如果父节点是红色那么一定存在爷爷节点，因为root节点不可能为空
		var uncle *rbNode = nil
		if parent == gParent.left { //左子节点的情况
			uncle = gParent.right
			if uncle != nil && r.isRed(uncle) { // 父亲和叔叔都是红色
				// 将父亲和叔叔染色为黑色，爷爷染色为红色
				r.setBlack(parent)
				r.setBlack(uncle)
				r.setRed(gParent)
				r.reBalance(gParent) //以爷爷节点为当前节点进行处理，进行以爷爷为当前节点进行旋转或其他操作
				return
			}
			// 1.2、叔叔节点不存在 | 或者为黑色，父亲节点为爷爷节点的左子树
			if uncle == nil || r.isBlack(uncle) {
				// 1.2.1、插入节点为父亲节点的左子节点（LL）将父亲染为黑色，将爷爷染为红色，然后以爷爷节点右旋，完成
				if node == parent.left {
					r.setBlack(parent)
					r.setRed(gParent)
					r.rightRotate(gParent)
					return
				}
				// 1.2.2、插入节点为父亲节点的右子节点（LR）以父亲节点进行一次左旋，得到（LL）然后按照1.2.1处理
				if node == parent.right {
					r.leftRotate(parent)
					r.reBalance(parent) //为什么要以父亲节点来重新处理呢？因为进行左旋后父亲节点被换到叶子节点去了
					return
				}
			}
		} else {
			uncle = gParent.left
			if uncle != nil && r.isRed(uncle) { //叔父双红
				// 将父亲和叔叔染色为黑色，爷爷染色为红色
				r.setBlack(parent)
				r.setBlack(uncle)
				r.setRed(gParent)
				r.reBalance(gParent) //以爷爷节点为当前节点进行处理，进行以爷爷为当前节点进行旋转或其他操作
				return
			}
			// 1.3、叔叔节点不存在 | 或者为黑色，父亲节点为爷爷节点的右子树
			if uncle == nil || r.isBlack(uncle) {
				// 1.3.1、插入的节点为父亲节点的右子树（RR），将父亲染为黑色，爷爷染为红色，最后以爷爷为节点进行左旋，完成
				if node == parent.right {
					r.setBlack(parent)
					r.setRed(gParent)
					r.leftRotate(gParent)
					return
				}
				// 1.3.2、插入的节点为父亲节点的左子树（LR），以父亲节点进行一次右旋，得到（RR），然后按照1.3.1处理
				if node == parent.left {
					r.rightRotate(parent)
					r.reBalance(parent)
					return
				}
			}
		}
	}
}

//
// leftRotate
// 树的左旋操作
//
//        p				p
//        |				|
//  	  x				y
// 		 / \	 ->	   / \
//      xl  y   	  x   yr
//         / \       / \
//        yl  yr    xl  yl
//
func (r *RBTree) leftRotate(x *rbNode) {
	var y = x.right
	x.right = y.left
	// 1、将x的右子节点改为y的左子节点
	if y.left != nil {
		y.left.parent = x
	}
	// 2、当x的parent不为空时，更新y的parent，和p的left和p的right为当前的y
	if x.parent != nil {
		y.parent = x.parent
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	} else {
		y.parent = nil
		r.Root = y
	}
	// 3、将x的parent设置为y，y的left设置为x
	x.parent = y
	y.left = x
}

//
// rightRotate
// 树的右旋操作
//
//        p				p
//        |				|
//  	  x				y
// 		 / \	 ->	   / \
//      y   xr   	  yl  x
//     / \               / \
//    yl  yr            yr  xr
//
func (r *RBTree) rightRotate(x *rbNode) {
	var y = x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	if x.parent != nil {
		y.parent = x.parent
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	} else {
		y.parent = nil
		r.Root = y
	}
	x.parent = y
	y.right = x
}

//
// PrintRBTree
// 二叉树打印（只为美观）
// 该代码抄自C++版本，(-_-)感觉我已经入门C++了呀
// 参考地址：https://www.tqwba.com/x_d/jishu/288287.html
//
func PrintRBTree(root *rbNode) {
	if root == nil {
		return
	}
	var tmpNode = root
	var intv = tmpNode.InorderTraversal()
	var tmp_str = ""
	var location = 0
	var recordFirst = make(map[*rbNode]int)
	for _, val := range intv {
		location = len(tmp_str)
		tmp_str += strconv.Itoa(val.key) + " "
		recordFirst[val] = location
	}
	var tmpstrLen = len(tmp_str)
	for i := 0; i < tmpstrLen; i++ { //全部置成 " "
		tmp_str = tmp_str[:i] + " " + tmp_str[i+1:]
	}
	var queue = make([]*rbNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		var currentSize = len(queue)
		var cur_loc = 0
		var tmpStr1 = tmp_str
		var tmpStr2 = tmp_str
		for i := 1; i <= currentSize; i++ {
			var node = queue[0]
			queue = append(queue[:0], queue[1:]...)
			cur_loc = recordFirst[node]
			var num_str = strconv.Itoa(node.key)
			if node.left != nil {
				queue = append(queue, node.left)
				var first_loc = recordFirst[node.left] + 1
				tmpStr2 = tmpStr2[:first_loc] + "/" + tmpStr2[first_loc+1:]
				first_loc++
				for first_loc < cur_loc {
					tmpStr1 = tmpStr1[:first_loc] + "_" + tmpStr1[first_loc+1:]
					first_loc++
				}
			}
			for j := 0; j < len(num_str); j++ {
				tmpStr1 = tmpStr1[:cur_loc] + num_str[j:j+1] + tmpStr1[cur_loc+1:]
				cur_loc++
			}
			if node.right != nil {
				queue = append(queue, node.right)
				var last_loc = recordFirst[node.right] - 1
				tmpStr2 = tmpStr2[:last_loc] + "\\" + tmpStr2[last_loc+1:]
				for cur_loc < last_loc {
					tmpStr1 = tmpStr1[:cur_loc] + "_" + tmpStr1[cur_loc+1:]
					cur_loc++
				}
			}
		}
		fmt.Println(tmpStr1)
		fmt.Println(tmpStr2)
	}
}
