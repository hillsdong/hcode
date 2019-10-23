package list

import "testing"

func TestList(t *testing.T) {
	l := New()
	t.Log(l)

	n := l.Header()
	n = l.InsertValue(1, n)
	n = l.InsertValue(2, n)
	n = l.InsertValue(3, n)
	n4 := l.InsertValue(4, n)
	l.InsertValue(5, n4)
	t.Log(l)

	l.Remove(n4)
	t.Log(l)

	l.Reverse()
	t.Log(l)
}
