package sort

import "testing"

func TestMaxK(t *testing.T) {
	nums := []int{2, 2, 1}
	t.Log(MaxK(nums, 2))
}
