package leetcode

// hash
func hasCycle(head *ListNode) bool {
	hash := map[*ListNode]bool{}
	for head != nil {
		if hash[head] {
			return true
		}
		hash[head] = true
		head = head.Next
	}
	return false
}

// 双指针
func hasCycle2(head *ListNode) bool {
	i, j := head, head
	for j != nil && j.Next != nil {
		i = i.Next
		j = j.Next.Next
		if j == i {
			return true
		}
	}
	return false
}
