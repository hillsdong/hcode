package leetcode

// 双指针
func detectCycle2(head *ListNode) *ListNode {
	i, j := head, head
	for j != nil && j.Next != nil {
		i = i.Next
		j = j.Next.Next
		if j == i {
			i = head
			for j != i {
				i = i.Next
				j = j.Next
			}
			return j
		}
	}
	return nil
}

// 哈希
func detectCycle(head *ListNode) *ListNode {
	hash := map[*ListNode]bool{}
	for head != nil {
		if hash[head] {
			return head
		}
		hash[head] = true
		head = head.Next
	}
	return nil
}

// func detectCycle2(head *ListNode) *ListNode {
// 	i, j := head, head
// 	in, jn := 0
// 	for j != nil && j.Next != nil {
// 		i = i.Next
// 		j = j.Next.Next
// 		in += 1
// 		jn += 2
// 		if j == i {
// 			break
// 		}
// 	}
// 	k := head
// 	for i := 0; i <= in; i++ {
// 		l = in - 1 + i
// 		if l - i + 1 = in
// 	}
// 	return head
// }
