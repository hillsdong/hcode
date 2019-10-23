package lineartable

import "testing"

func TestArray(t *testing.T) {
	a := NewArray(5)
	a.Add(0, 0)
	t.Log(a)

	for i := 1; i < 6; i++ {
		a.Add(i, i)
		t.Log(a)
	}

	a.Remove(5)
	t.Log(a)

	a.Set(3, 44)
	a.Set(4, 44)
	t.Log(a)

	t.Log(a.Get(4))
	t.Log(a.Get(1))

	t.Log(a.Find(44))
	t.Log(a.FindAll(44))
}
