package main

import "fmt"

func main() {
	var arr = []int{100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0}
	var res = distributeCandies(arr)
	fmt.Println(res)
}

func distributeCandies(candyType []int) int {
	var arrLen = len(candyType)
	if arrLen == 0 || arrLen == 1 {
		return arrLen
	}
	var record = make(map[int]int)
	var canEat = arrLen / 2
	var maxVal = 0
	for _, val := range candyType {
		if record[val] == 0 {
			record[val] = 1
			maxVal++
			if maxVal >= canEat {
				return canEat
			}
		}
	}
	return maxVal
}
