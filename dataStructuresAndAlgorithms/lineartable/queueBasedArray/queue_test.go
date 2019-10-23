package queue

import "testing"

func TestQueue(t *testing.T) {
	s := New(5)
	t.Log(s.Push(1))
	t.Log(s.Push(2))
	t.Log(s.Push(3))
	t.Log(s.Push(4))
	t.Log(s)
	t.Log(s.Push(5))
	t.Log(s)
	t.Log(s.Push(6))
	t.Log(s)

	t.Log(s.Pull())
	t.Log(s.Pull())
	t.Log(s.Pull())
	t.Log(s.Pull())
	t.Log(s)
	t.Log(s.Pull())
	t.Log(s.Pull())
	t.Log(s)

	t.Log(s.Push(1))
	t.Log(s.Push(2))
	t.Log(s.Push(3))
	t.Log(s.Push(4))
	t.Log(s)

	t.Log(s.Pull())
	t.Log(s.Pull())
	t.Log(s)

}
