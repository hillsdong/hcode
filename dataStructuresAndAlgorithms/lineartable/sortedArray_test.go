package lineartable

import "testing"

func TestSortedArray(t *testing.T) {
	a := NewSortedArray(5)

	nums := [5]int64{23, 1, 124, 56, 7}

	for _, n := range nums {
		a.Add(n)
		t.Log(a)
	}

	a.Remove(23)
	t.Log(a)

	t.Log(a.Find(124))

}
