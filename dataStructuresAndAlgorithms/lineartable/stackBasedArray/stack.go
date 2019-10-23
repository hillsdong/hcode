package stack

import (
	"bytes"
	"errors"
	"fmt"
)

// New a Stack
func New(n int) *Stack {
	return &Stack{make([]interface{}, n, n), -1, n}
}

// Stack struct
type Stack struct {
	data []interface{}
	top  int
	cap  int
}

// Push pushes a value to stack, return full error or nil
func (s *Stack) Push(v interface{}) error {
	if s.top >= s.cap-1 {
		return errors.New("the stack is full")
	}
	s.top++
	s.data[s.top] = v
	return nil
}

// Pop popes a value from stack, return value and error
func (s *Stack) Pop() (interface{}, error) {
	if s.top < 0 {
		return nil, errors.New("the stack is empty")
	}
	defer func() { s.top-- }()
	return s.data[s.top], nil
}

func (s *Stack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Stack: length = %d, capacity = %d, ", s.top+1, s.cap))
	buffer.WriteString("data = [")
	for i := 0; i <= s.top; i++ {
		buffer.WriteString(fmt.Sprint(s.data[i]))
		if i < s.top {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(fmt.Sprintf("]"))
	return buffer.String()
}
