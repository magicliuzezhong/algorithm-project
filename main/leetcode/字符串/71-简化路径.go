//
// Package main
// @Description：
// @Author：liuzezhong 2021/9/18 15:53
// @Company cloud-ark.com
//
package main

import (
	"fmt"
	"strings"
)

func main() {
	var paths = []string{
		//"/home/",
		//"/../",
		//"/home//foo/",
		"/a/./b/../../c/",
		"/home/././abc//bb/a/../../",
	}
	for _, path := range paths {
		var result = simplifyPath(path)
		fmt.Println(result)
	}
}

//
// simplifyPath
// 第一时间就想到用栈实现了，事实确实如此
// 测试发现
// strings.Join()函数使得执行时间从0ms变动到4ms，但是内存消耗从4.2MB变为了3.1MB，
// 感觉还是手工拼接比strings.Join()好一些，用空间换时间还是比较划算的
//
func simplifyPath(path string) string {
	// /home/./../abc//bb/a/../../ /
	var paths = strings.Split(path, "/")
	var stack = make([]string, 0)
	for _, val := range paths {
		if len(val) == 0 || val == "." {
			continue
		}
		if val == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] //出栈
			}
		} else {
			stack = append(stack, val)
		}
	}
	// 这种方式比较推荐的，原因在函数注释上
	//var result = ""
	//for len(stack) > 0 {
	//	result = "/" + stack[len(stack)-1] + result
	//	stack = stack[:len(stack)-1]
	//}
	//if len(result) == 0 {
	//	return "/"
	//}
	//return result
	// 还是不推荐使用这个函数，当然项目上就无所谓了，能少写代码还是很棒的
	return "/" + strings.Join(stack, "/")
}
