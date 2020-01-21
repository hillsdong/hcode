package bag01

import "testing"

func TestBag01(t *testing.T) {
	b := New([]int{2, 2, 3, 3}, 9)
	t.Log(b.MaxWeight())
	t.Log(b.status)
}
