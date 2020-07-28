package leetcode

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	l := &ListNode{}
	h := l
	for i := 1; i < 8; i++ {
		l.Val, l.Next = i, &ListNode{Val: i + 1}
		l = l.Next
	}

	l = swapPairs(h)
	fmt.Println(l)
}

func TestSwapPairs2(t *testing.T) {
	l := &ListNode{}
	h := l
	for i := 1; i < 2; i++ {
		l.Val, l.Next = i, &ListNode{Val: i + 1}
		l = l.Next
	}

	l = swapPairs(h)
	fmt.Println(l)
}

func TestSwapPairs3(t *testing.T) {
	l := &ListNode{Val: 1}
	h := l

	l = swapPairs(h)
	fmt.Println(l)
}
