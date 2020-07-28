package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归 简洁
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := reverseList(head.Next)
	head.Next.Next, head.Next = head, nil
	return tail
}

// 递归
func reverseList4(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	return _reverseList(pre, cur)
}

func _reverseList(pre, cur *ListNode) *ListNode {
	for cur == nil {
		return pre
	}
	pre, cur, cur.Next = cur, cur.Next, pre
	return _reverseList(pre, cur)
}

// 循环
func reverseList3(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

// 循环 简化
func reverseList2(head *ListNode) (pre *ListNode) {
	for cur := head; cur != nil; pre, cur, cur.Next = cur, cur.Next, pre {
	}
	return
}
