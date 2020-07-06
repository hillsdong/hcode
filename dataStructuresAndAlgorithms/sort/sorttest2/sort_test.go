package sorttest2

import "testing"

func TestBubble(t *testing.T) {
	nums := []int{3, 2, 4, 1, 6, 6, 5}
	t.Log(Bubble(nums))
}

func TestInsertion(t *testing.T) {
	nums := []int{3, 2, 4, 1, 6, 6, 5}
	t.Log(Insertion(nums))
}

func TestSelectoin(t *testing.T) {
	nums := []int{3, 2, 4, 1, 6, 6, 5}
	t.Log(Selection(nums))
}

func TestMerge(t *testing.T) {
	nums := []int{3, 2, 4, 1, 6, 6, 5}
	t.Log(Merge(nums))
}

func TestQuick(t *testing.T) {
	nums := []int{3, 2, 4, 1, 6, 6, 5}
	t.Log(Quick(nums))
}
