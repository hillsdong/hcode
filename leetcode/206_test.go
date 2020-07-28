package leetcode

import "testing"

func TestReverseList(t *testing.T) {
	l := &ListNode{}
	h := l
	for i := 1; i < 6; i++ {
		l.Val, l.Next = i, &ListNode{}
		l = l.Next
	}
	l.Next = nil

	l = reverseList(h)
	for l.Next != nil {
		l = l.Next
		t.Log(l.Val)
	}
}
