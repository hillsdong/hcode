package queue

import (
	"bytes"
	"errors"
	"fmt"
)

// New a circular queue, no data at head
func New(n int) *Queue {
	return &Queue{make([]interface{}, n+1, n+1), 0, 0, n + 1}
}

// Queue struct
type Queue struct {
	data []interface{}
	head int
	tail int
	cap  int
}

// Push pushes a value to queue's tail, return full error or nil
func (q *Queue) Push(v interface{}) error {
	if (q.tail+1)%q.cap == q.head {
		return errors.New("the queue is full")
	}
	q.tail = (q.tail + 1) % q.cap
	q.data[q.tail] = v
	return nil
}

// Pull pulles a value from queue's head, return value and error
func (q *Queue) Pull() (interface{}, error) {
	if q.head == q.tail {
		return nil, errors.New("the queue is empty")
	}
	q.head = (q.head + 1) % q.cap
	return q.data[q.head], nil
}

func (q *Queue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Queue: length = %d, capacity = %d, head = %d, tail = %d, ", (q.tail+q.cap-q.head)%q.cap, q.cap-1, q.head, q.tail))
	buffer.WriteString("data = [")
	tail := q.tail
	if tail < q.head {
		tail = tail + q.cap
	}
	for i := q.head + 1; i <= tail; i++ {
		buffer.WriteString(fmt.Sprint(q.data[i%q.cap]))
		if i != tail {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")

	buffer.WriteString("array = [")
	for i := 0; i < q.cap; i++ {
		buffer.WriteString(fmt.Sprint(q.data[i]))
		if i != q.cap-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
