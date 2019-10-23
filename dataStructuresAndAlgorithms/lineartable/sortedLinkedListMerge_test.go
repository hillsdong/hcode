package lineartable

import "testing"

func TestSortedLinkedListMerge(t *testing.T) {
	a := NewLinkedList()
	a.Insert(1, 1)
	a.Insert(2, 3)
	a.Insert(3, 5)

	b := NewLinkedList()
	b.Insert(1, 2)
	b.Insert(2, 4)
	b.Insert(3, 6)

	t.Log(SortedLinkedListMerge(NewLinkedList(), b))
	t.Log(SortedLinkedListMerge(a, NewLinkedList()))
	t.Log(SortedLinkedListMerge(NewLinkedList(), NewLinkedList()))
	t.Log(SortedLinkedListMerge(a, b))
}
