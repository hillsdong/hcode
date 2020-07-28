package leetcode

import "fmt"

func longestCommonSubsequence(text1, text2 string) int {
	l1, l2 := len(text1), len(text2)
	c := make([]int, l1+1)
	var _jpre, jpre int
	for i := l2 - 1; i >= 0; i-- {
		for j := l1 - 1; j >= 0; j-- {
			_jpre = c[j]
			if text2[i] == text1[j] {
				c[j] = jpre + 1
			} else {
				c[j] = max(c[j], c[j+1])
			}
			jpre = _jpre
		}
		fmt.Println(c)
		jpre = 0
	}
	return c[0]
}
