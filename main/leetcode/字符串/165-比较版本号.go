//
// Package 字符串
// @Description：
// @Author：liuzezhong 2021/9/16 09:10
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var version1 = "1.0.0"
	var version2 = "1"
	var result = compareVersion(version1, version2)
	fmt.Println(result)
}

func compareVersion(version1 string, version2 string) int {
	var v1 = strings.Split(version1, ".")
	var v2 = strings.Split(version2, ".")
	var v1Len = len(v1)
	var v2Len = len(v2)
	var i = 0
	for i < v1Len && i < v2Len {
		var v1Num, _ = strconv.Atoi(v1[i])
		var v2Num, _ = strconv.Atoi(v2[i])
		if v1Num > v2Num {
			return 1
		} else if v1Num < v2Num {
			return -1
		}
		i++
	}
	for i < v1Len {
		var v1Num, _ = strconv.Atoi(v1[i])
		if v1Num > 0 {
			return 1
		}
		i++
	}
	for i < v2Len {
		var v2Num, _ = strconv.Atoi(v2[i])
		if v2Num > 0 {
			return -1
		}
		i++
	}
	return 0
}
