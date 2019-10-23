package lineartable

import "testing"

func TestSortedArrayMerge(t *testing.T) {
	a := []int64{1, 3, 5, 7, 9, 10, 11}
	b := []int64{2, 4, 6, 8}
	t.Log(SortedArrayMerge([]int64{}, b))
	t.Log(SortedArrayMerge(a, []int64{}))
	t.Log(SortedArrayMerge([]int64{}, []int64{}))
	t.Log(SortedArrayMerge(a, b))
}

func TestSortedArrayMerge2(t *testing.T) {
	a := []int64{1, 3, 5, 7, 9, 10, 11}
	b := []int64{2, 4, 6, 8}
	t.Log(SortedArrayMerge2([]int64{}, b))
	t.Log(SortedArrayMerge2(a, []int64{}))
	t.Log(SortedArrayMerge2([]int64{}, []int64{}))
	t.Log(SortedArrayMerge2(a, b))
}
