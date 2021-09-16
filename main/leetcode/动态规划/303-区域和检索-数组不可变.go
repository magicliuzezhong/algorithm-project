//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/16 08:53
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var arr = []int{-2, 0, 3, -5, 2, -1}
	var numArray = Constructor1(arr)
	var result1 = numArray.SumRange(0, 2) // return 1 ((-2) + 0 + 3)
	var result2 = numArray.SumRange(2, 5) // return -1 (3 + (-5) + 2 + (-1))
	var result3 = numArray.SumRange(0, 5) // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

type NumArray struct {
	sum []int
}

func Constructor1(nums []int) NumArray {
	var numLen = len(nums)
	var instance = NumArray{}
	instance.sum = make([]int, numLen+1)
	for i := 1; i <= numLen; i++ {
		instance.sum[i] = instance.sum[i-1] + nums[i-1]
	}
	return instance
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum[right+1] - this.sum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
