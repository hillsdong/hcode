package leetcode

import "testing"

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	t.Log(minimumTotal(triangle))
}
