package leetcode

import "testing"

func TestMinStack(t *testing.T) {
	minStack := ConstructorMinStack2()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	t.Log(minStack.GetMin()) //   --> 返回 -3.
	minStack.Pop()
	t.Log(minStack.Top())    //   --> 返回 0.
	t.Log(minStack.GetMin()) //   --> 返回 -2.

}
