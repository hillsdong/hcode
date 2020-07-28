package leetcode

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{3, 0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}
