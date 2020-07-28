package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	l := &ListNode{}
	h := l
	for i := 1; i < 8; i++ {
		l.Val, l.Next = i, &ListNode{Val: i + 1}
		l = l.Next
	}

	l = reverseKGroup(h, 3)
	fmt.Println(l)
}

func TestReverseKGroup2(t *testing.T) {
	l := &ListNode{}
	h := l
	for i := 1; i < 6; i++ {
		l.Val, l.Next = i, &ListNode{Val: i + 1}
		l = l.Next
	}

	l = reverseKGroup(h, 3)
	fmt.Println(l)
}

func TestReverseKGroup3(t *testing.T) {
	l := &ListNode{}
	h := l
	for i := 1; i < 3; i++ {
		l.Val, l.Next = i, &ListNode{Val: i + 1}
		l = l.Next
	}

	l = reverseKGroup(h, 3)
	fmt.Println(l)
}

func TestReverseKGroup4(t *testing.T) {
	l := &ListNode{}
	h := l
	for i := 1; i < 2; i++ {
		l.Val, l.Next = i, &ListNode{Val: i + 1}
		l = l.Next
	}

	l = reverseKGroup(h, 3)
	fmt.Println(l)
}

func TestReverseKGroup5(t *testing.T) {
	l := &ListNode{Val: 1}
	h := l

	l = reverseKGroup(h, 3)
	fmt.Println(l)
}

func TestReverseKGroup6(t *testing.T) {
	l := &ListNode{}
	h := l
	l = reverseKGroup(h, 3)
	fmt.Println(l)
}
