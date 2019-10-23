package stack

import (
	"bytes"
	"errors"
	"fmt"
)

// New a Stack
func New() *Stack {
	return &Stack{}
}

// Node defines a stack node
type Node struct {
	data interface{}
	next *Node
}

// Stack struct
type Stack struct {
	top *Node
	len int
}

// Push pushes a value to stack, return full error or nil
func (s *Stack) Push(v interface{}) error {
	s.top = &Node{v, s.top}
	s.len++
	return nil
}

// Pop popes a value from stack, return value and error
func (s *Stack) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("the stack is empty")
	}
	defer func() { s.top = s.top.next; s.len-- }()
	return s.top.data, nil
}

// Len return the len
func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("StackBasedLinkedList: length = %d, capacity = ~, ", s.len))
	buffer.WriteString("data = [")
	for n := s.top; n != nil; n = n.next {
		buffer.WriteString(fmt.Sprint(n.data))
		if n.next != nil {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(fmt.Sprintf("]"))
	return buffer.String()
}
