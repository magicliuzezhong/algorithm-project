package main

import "fmt"

func main() {
	//var s = "aaaaaaabc"
	//var goal = "aaaaaaacb"
	var s = "aa"
	var goal = "aa"
	var res = buddyStrings(s, goal)
	fmt.Println(res)
}

func buddyStrings(s string, goal string) bool {
	var sLen = len(s)
	var goalLen = len(goal)
	if sLen != goalLen {
		return false
	}
	var site1 = -1
	var site2 = -1
	var record = [26]int{}
	var hasDuplicate = false
	var count = 0
	for i := 0; i < sLen; i++ {
		if s[i] != goal[i] {
			if count == 0 {
				site1 = i
			} else {
				site2 = i
			}
			count++
			if count > 2 {
				return false
			}
		}
		if record[s[i]-'a'] == 0 {
			record[s[i]-'a'] = i + 1
		} else {
			hasDuplicate = true
		}
	}
	return count == 0 && hasDuplicate || count == 2 && s[site1] == goal[site2] && s[site2] == goal[site1]
}
