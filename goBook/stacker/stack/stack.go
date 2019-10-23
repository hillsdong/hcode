package stack

import "errors"

//Stack is a good type
type Stack []interface{}

//Push a value
func (stack *Stack) Push(val interface{}) error {
	*stack = append(*stack, val)
	return nil
}

//Pop a value
func (stack *Stack) Pop() (interface{}, error) {
	thestack := *stack
	if len(thestack) == 0 {
		return nil, errors.New("nothing in this stack")
	}
	val := thestack[len(thestack)-1]
	*stack = thestack[:len(thestack)-1]
	return val, nil
}

//Len is good function
func (stack Stack) Len() int {
	return len(stack)
}

//Top is good function
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("nothing in this stack")
	}
	return stack[len(stack)-1], nil
}
