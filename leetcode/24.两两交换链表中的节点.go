package leetcode

// TODO 递归

// 迭代 哨兵优化
func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	for p := pre; p.Next != nil && p.Next.Next != nil; p = p.Next.Next {
		f := p.Next
		s := f.Next

		f.Next, s.Next = s.Next, f
		// 关键
		p.Next = s
	}
	return pre.Next
}

// 迭代
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p, f, s *ListNode = nil, head, head.Next
	var ret *ListNode = head
	for f != nil && s != nil {
		f.Next, s.Next = s.Next, f

		if p != nil {
			p.Next = s
		} else {
			ret = s
		}

		if f.Next == nil {
			break
		}

		p, f, s = f, f.Next, f.Next.Next
	}
	return ret
}
