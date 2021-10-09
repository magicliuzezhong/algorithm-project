//
// Package main
// @Description：
// @Author：liuzezhong 2021/10/9 10:48
// @Company cloud-ark.com
//
package main

import "fmt"

func main() {
	var result = firstBadVersion(1)
	fmt.Println(result)
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	if version == 1 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	var left = 1
	var right = n
	var mid = 0
	for left <= right {
		mid = (left + right) / 2
		if (mid == 0 && isBadVersion(mid)) || (!isBadVersion(mid-1) && isBadVersion(mid)) {
			return mid
		}
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
