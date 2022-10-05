// Package main
// Author 刘泽忠
// Time 2022-10-05 21:32:23
// Description:
// Copyright: 2018-2022 常相伴（武汉）科技有限公司 版权所有
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var arr = []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	res := subdomainVisits(arr)
	fmt.Println(res)
}

func subdomainVisits(cpdomains []string) []string {
	counts := make(map[string]int)
	for _, val := range cpdomains {
		tmpSplit := strings.Split(val, " ")
		count, _ := strconv.Atoi(tmpSplit[0])
		word := tmpSplit[1]
		counts[word] += count
		for i, ch := range word {
			if ch == '.' {
				counts[word[i+1:]] += count
			}
		}
	}

	res := make([]string, len(counts))
	index := 0
	for word, count := range counts {
		res[index] = fmt.Sprintf("%d %s", count, word)
		index++
	}
	return res
}
