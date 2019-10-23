package list

import (
	"bytes"
	"fmt"
)

// Node is a node of a doubly linked list.
type Node struct {
	prev, next *Node
	Data       interface{}
}

// Prev returns the prev node.
func (n *Node) Prev() *Node {
	return n.prev
}

// Next returns the next node.
func (n *Node) Next() *Node {
	return n.next
}

// List is a doubly linked list.
type List struct {
	root *Node
	len  int
}

// New returns an doubly linked List.
func New() *List {
	root := &Node{}
	root.prev, root.next = root, root
	return &List{root, 0}
}

// Len returns the count of nodes but not include the root.
func (l *List) Len() int { return l.len }

// Header returns the first node.
func (l *List) Header() *Node {
	return l.root.next
}

// Tailer returns the last node.
func (l *List) Tailer() *Node {
	return l.root.prev
}

// insert inserts n after at, increments len, and returns n.
func (l *List) insert(n, front *Node) *Node {
	back := front.next
	n.prev, n.next = front, back
	front.next, back.prev = n, n
	l.len++
	return n
}

// InsertValue is a wrapper for inert(&Node{Data: v}, front).
func (l *List) InsertValue(v interface{}, front *Node) *Node {
	return l.insert(&Node{Data: v}, front)
}

// Remove remove n, decrements len, and returns n.
func (l *List) Remove(n *Node) *Node {
	n.prev.next, n.next.prev = n.next, n.prev
	n.prev, n.next = nil, nil // avoid memory leaks
	l.len--
	return n
}

// Reverse reverses the list
func (l *List) Reverse() {
	l.root.prev, l.root.next = l.root.next, l.root.prev
	for node := l.Header(); node != l.root; node = node.Next() {
		node.prev, node.next = node.next, node.prev
	}
}

func (l *List) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Doubly Liked List: Length = %d \n", l.Len()))
	buffer.WriteString(fmt.Sprintf("Root = %p, %d, %p, %p -> \n", l.root, l.root.Data, l.root.prev, l.root.next))
	if l.Len() != 0 {
		for node := l.Header(); node != l.root; node = node.Next() {
			buffer.WriteString(fmt.Sprintf("Data = %p, %d, %p, %p -> \n", node, node.Data, node.prev, node.next))
		}
	}
	return buffer.String()
}
