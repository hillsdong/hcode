package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	t.Log(s.Push(1))
	t.Log(s.Push(2))
	t.Log(s.Push(3))
	t.Log(s.Push(4))
	t.Log(s)
	t.Log(s.Push(5))
	t.Log(s.Push(6))
	t.Log(s)

	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s)
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s)

}
