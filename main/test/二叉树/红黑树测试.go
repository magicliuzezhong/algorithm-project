//
// Package main
// @Description：
// @Author：liuzezhong 2021/8/30 17:25
// @Company cloud-ark.com
//
package main

import (
	. "../../../model"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var tree = &RBTree{}
	var input = bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var line = input.Text()
		var key, _ = strconv.Atoi(line)
		tree.Insert(key, "adsdas")
		PrintRBTree(tree.Root)
	}
}
