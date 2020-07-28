package leetcode

import "testing"

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	k := 3
	t.Log("ramdom")
	t.Log(topKFrequent(nums, k))
}

type IntQS []int

func (is IntQS) Len() int {
	return len(is)
}

func (is IntQS) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntQS) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
	return
}

func TestQuickSelect(t *testing.T) {
	nums := IntQS{1, 4, 5, 0, 2, 3}
	k := 3
	quickSelect(nums, k)
	t.Log(nums[len(nums)-k:])
}
