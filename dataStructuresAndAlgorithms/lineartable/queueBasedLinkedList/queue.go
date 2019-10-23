package queue

import (
	"bytes"
	"errors"
	"fmt"
)

// New a circular queue
func New() *Queue {
	n := &Node{}
	n.next = n
	return &Queue{n, n, 0}
}

// Node defines a queue node
type Node struct {
	data interface{}
	next *Node
}

// Queue struct
type Queue struct {
	head *Node
	tail *Node
	len  int
}

// Push pushes a value to queue's tail, return full error or nil
func (q *Queue) Push(v interface{}) error {
	n := &Node{v, q.head}

	q.tail.next = n
	q.tail = n

	q.len++
	return nil
}

// Pull pulles a value from queue's head, return value and error
func (q *Queue) Pull() (interface{}, error) {
	if q.head == q.tail {
		return nil, errors.New("the queue is empty")
	}
	defer func() {
		q.head = q.head.next
		q.tail.next = q.head
		q.len--
	}()
	return q.head.next, nil
}

func (q *Queue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Queue: length = %d, head = %v, tail = %v, ", q.len, q.head, q.tail))
	buffer.WriteString("data = [")
	for n := q.head.next; n != q.head; n = n.next {
		buffer.WriteString(fmt.Sprint(n.data))
		if n.next != q.head {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
