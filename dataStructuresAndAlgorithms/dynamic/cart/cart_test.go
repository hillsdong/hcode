package cart

import "testing"

func TestCart(t *testing.T) {
	c := New([]int{1, 2, 4, 3}, 9)
	t.Log(c.MinPrice())
	t.Log(c.status)
	t.Log(c.MinPriceItems())
}
