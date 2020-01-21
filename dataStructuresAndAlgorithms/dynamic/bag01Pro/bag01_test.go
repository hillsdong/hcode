package bag01

import "testing"

func TestBag01(t *testing.T) {
	b := New([]int{2, 2, 3, 3, 10}, []int{20, 2, 3, 3, 30}, 9)
	t.Log(b.MaxValue())
	t.Log(b.status)
}
