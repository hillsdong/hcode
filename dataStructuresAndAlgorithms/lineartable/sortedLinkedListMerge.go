package lineartable

// SortedLinkedListMerge merges two sorted linked list
func SortedLinkedListMerge(a, b *LinkedList) *LinkedList {
	if a.Len == 0 {
		return b
	}
	if b.Len == 0 {
		return a
	}
	c := NewLinkedList()
	an, bn, cn := a.Head.Next, b.Head.Next, c.Head

	for an != nil && bn != nil {
		if an.Data.(int) < bn.Data.(int) {
			cn.Next = an
			cn = cn.Next
			an = an.Next
		} else {
			cn.Next = bn
			cn = cn.Next
			bn = bn.Next
		}
	}

	for an != nil {
		cn.Next = an
		an = an.Next
	}

	for bn != nil {
		cn.Next = bn
		bn = bn.Next
	}
	return c
}
