package leetcode

import "testing"

func TestHasCycle(t *testing.T) {
	l := &ListNode{}
	h := l
	for i := 1; i < 6; i++ {
		l.Val, l.Next = i, &ListNode{}
		l = l.Next
	}
	l.Next = h
	t.Log(hasCycle(h))
}
