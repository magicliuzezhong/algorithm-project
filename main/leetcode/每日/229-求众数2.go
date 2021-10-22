//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/22 16:23
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{1, 1, 1, 3, 3, 2, 2, 2}
	var result = majorityElement(arr)
	var result1 = majorityElement1(arr)
	fmt.Println(result)
	fmt.Println(result1)
}

//
// majorityElement
// 怎么在O(n)的时间复杂度O(1)空间复杂度解决才是困难点？
// 说明只能一次遍历，外加不能使用过长的辅助字段
// 看了下题解，摩尔投票法，我和我的小伙伴儿都惊呆了
// 类似于消消乐，但是消除的是不相同的，相同的就相加
//
func majorityElement(nums []int) []int {
	var ele1, ele2, vote1, vote2, factor = 0, 0, 0, 0, len(nums) / 3
	for _, val := range nums {
		if vote1 > 0 && ele1 == val {
			vote1++
		} else if vote2 > 0 && ele2 == val {
			vote2++
		} else if vote1 == 0 {
			ele1 = val
			vote1 = 1
		} else if vote2 == 0 {
			ele2 = val
			vote2 = 1
		} else {
			vote1--
			vote2--
		}
	}
	vote1, vote2 = 0, 0
	for _, val := range nums {
		if val == ele1 {
			vote1++
		} else if val == ele2 {
			vote2++
		}
	}
	var result []int
	if vote1 > factor {
		result = append(result, ele1)
	}
	if vote2 > factor {
		result = append(result, ele2)
	}
	return result
}

//
// majorityElement
// @Description: 暴力破解
//
func majorityElement1(nums []int) (result []int) {
	var factor = len(nums) / 3
	var record = make(map[int]int)
	for _, val := range nums {
		record[val]++
	}
	for key, val := range record {
		if val > factor {
			result = append(result, key)
		}
	}
	return
}
