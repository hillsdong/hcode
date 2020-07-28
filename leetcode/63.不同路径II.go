package leetcode

import "fmt"

// TODO
func ljjs(table [][]bool) int {
	l := len(table)
	c := make([]int, l+1)
	for j := l - 1; j >= 0; j-- {
		for i := l - 1; i >= 0; i-- {
			if table[i][j] {
				c[i] = 0
				continue
			}
			if i == l-1 && j == l-1 {
				c[i] = 1
				continue
			}
			c[i] += c[i+1]
		}
		fmt.Println(c)
	}
	return c[0]
}
