package main

import "fmt"

func main() {
	var arr = []int{3, 2, 1, 1, 4}
	var result = canJump(arr)
	fmt.Println(result)
}

func canJump(nums []int) bool {
	var numLen = len(nums)
	var mostVal = 0
	for i := 0; i < numLen-1 && mostVal < numLen-1; i++ {
		if mostVal < i+nums[i] {
			mostVal = i + nums[i]
		}
		if i == mostVal {
			break
		}
	}
	return mostVal >= numLen-1
}
