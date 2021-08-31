//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/26 16:10
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"math"
)

func main() {
	//      北京  天津  郑州  济南  长沙  海南  美国
	// 北京  0    100  1200  -    -    -    -
	// 天津  100  0    900   300  -    -    -
	// 郑州  1200 900  0     400  500  -    -
	// 济南  -    300  400   0    1300 1400 -
	// 长沙  -    -    500   1300 0    1500 -
	// 海南  -    -    -     1400 1500 0    -
	// 美国  -    -    -     -    -    -    -
	var graph = [][]int{
		{0, 100, 1200, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64},
		{100, 0, 900, 300, math.MaxInt64, math.MaxInt64, math.MaxInt64},
		{1200, 900, 0, 400, 500, math.MaxInt64, math.MaxInt64},
		{math.MaxInt64, 300, 400, 0, 1300, 1400, math.MaxInt64},
		{math.MaxInt64, math.MaxInt64, 500, 1300, 0, 1500, math.MaxInt64},
		{math.MaxInt64, math.MaxInt64, math.MaxInt64, 1400, 1500, 0, math.MaxInt64},
		{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64},
	}
	var result = dijkstra(graph, 0)
	fmt.Println(result)
}

func dijkstra(graph [][]int, src int) []int {
	var graphLen = len(graph)
	var visited = make([]bool, graphLen)
	var dist = make([]int, graphLen)
	// 进行初始化
	for i := 0; i < graphLen; i++ {
		dist[i] = graph[src][i]
	}
	visited[src] = true
	for i := 0; i < graphLen; i++ {
		// 先找到中间节点
		var minDist = math.MaxInt64
		var middle = i
		for j := 0; j < graphLen; j++ {
			// 这里使用的贪心的思想，找到没有被访问过的最小的哪一个节点
			if !visited[j] && minDist > dist[j] {
				minDist = dist[j]
				middle = j
			}
		}
		// 上述已经找到中间节点了
		// 现在需要找到当前中间节点所连接的相邻的节点
		// 根据上面找到的未被访问的当前最小的节点
		// 然后把当前节点与未被访问的节点进行相加判断是否小于当前dist的最小值
		// -> min > min可达 + 可达节点的值
		for j := 0; j < graphLen; j++ {
			if !visited[j] && graph[middle][j] != math.MaxInt64 && dist[j] > graph[middle][j]+minDist {
				dist[j] = graph[middle][j] + minDist
			}
		}
		visited[middle] = true
	}
	return dist
}
