//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/26 11:42
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = [][]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{4, 8},
		{5, 6},
		{5, 7},
		{7, 8},
		{9, 10},
		{11, 13},
		{12, 14},
		{11, 15},
		{11, 17},
	}
	var result = testConnect(arr, 4, 5)
	fmt.Println(result)
}

func testConnect(arr [][]int, a, b int) bool {
	var max = 0
	for _, val := range arr {
		var ele1 = val[0]
		var ele2 = val[1]
		if ele1 > max {
			max = ele1
		}
		if ele2 > max {
			max = ele2
		}
	}
	var unionSet = make([]int, max+1) //并查集
	var rank = make([]int, max+1)     //并查集
	for i, _ := range unionSet {
		unionSet[i] = i
		rank[i] = 1
	}
	for _, val := range arr {
		union(unionSet, rank, val[0], val[1])
	}
	return find(unionSet, a) == find(unionSet, b)
}

func union(unionSet []int, rank []int, ele1, ele2 int) {
	var root1 = find(unionSet, ele1)
	var root2 = find(unionSet, ele2)
	if root1 == root2 { // 不执行合并
		return
	}
	if rank[root1] > rank[root2] {
		unionSet[root2] = root1
		rank[root1] += rank[root2]
	} else {
		unionSet[root1] = root2
		rank[root2] += rank[root1]
	}
}

func find(unionSet []int, element int) int {
	for unionSet[element] != element {
		element = unionSet[element]
	}
	return element
}
