package leetcode

import "testing"

func TestDetectCycle(t *testing.T) {
	l := &ListNode{}
	h := l
	for i := 1; i < 6; i++ {
		l.Val, l.Next = i, &ListNode{}
		l = l.Next
	}
	l.Next = h.Next

	s := detectCycle(h)
	if s == nil {
		t.Log("nil")
	} else {
		t.Log(s.Val)
	}
}
