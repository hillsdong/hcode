package lineartable

import "testing"

func TestLinkedList(t *testing.T) {
	l := NewLinkedList()
	t.Log(l.Insert(0, 0))
	t.Log(l.Insert(1, 1))
	t.Log(l.Insert(2, 2))
	t.Log(l.Insert(4, 4))

	l.Insert(3, 3)
	l.Insert(4, 4)
	l.Insert(1, 11)
	t.Log(l)

	l.Reverse()
	t.Log(l)

	t.Log(l.FindMidNode())
	t.Log(l.FindMidNodeWithoutLen())

	t.Log(l.Remove(0))
	t.Log(l.Remove(1))
	t.Log(l.Remove(5))
	t.Log(l.Remove(4))
	t.Log(l)

	t.Log(l.Get(0))
	t.Log(l.Get(1))
	t.Log(l.Get(2))
	t.Log(l.Get(3))
	t.Log(l.Get(4))

	t.Log(l.Set(0, 0))
	t.Log(l.Set(1, 11))
	t.Log(l.Set(3, 11))
	t.Log(l.Set(4, 44))
	t.Log(l)

	t.Log(l.Find(11))
	t.Log(l.FindAll(11))
}
