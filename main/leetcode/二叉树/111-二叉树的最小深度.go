//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/26 10:16
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"fmt"
	"math"
)

func main() {
	//var tree = GetMinDepthTree()
	var tree = GetMinDepthTree1()
	//var tree = CreateTree1()
	var result = minDepth(tree)
	fmt.Println(result)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var minData = math.MaxInt32
	if root.Left != nil {
		minData = int(math.Min(float64(minDepth(root.Left)), float64(minData)))
	}
	if root.Right != nil {
		minData = int(math.Min(float64(minDepth(root.Right)), float64(minData)))
	}
	return minData + 1
}
