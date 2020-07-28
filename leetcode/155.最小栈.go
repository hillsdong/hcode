package leetcode

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}
func (s *stack) Pop() {
	if len(*s) == 0 {
		return
	}
	*s = (*s)[:len(*s)-1]
}
func (s *stack) Top() int {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}
func (s *stack) Empty() bool {
	return len(*s) == 0
}

// 辅助栈
type MinStack2 struct {
	dataStack *stack
	minStack  *stack
}

/** initialize your data structure here. */
func ConstructorMinStack2() MinStack2 {
	return MinStack2{&stack{}, &stack{}}
}

func (this *MinStack2) Push(x int) {
	this.dataStack.Push(x)
	min := this.minStack.Top()
	if x < min {
		min = x
	}
	this.minStack.Push(min)
}

func (this *MinStack2) Pop() {
	this.dataStack.Pop()
	this.minStack.Pop()
}

func (this *MinStack2) Top() int {
	return this.dataStack.Top()
}

func (this *MinStack2) GetMin() int {
	return this.minStack.Top()
}

// 最小变量
type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{stack: []int{}}
}

func (this *MinStack) Push(x int) {
	if x < this.min || len(this.stack) == 0 {
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if this.Top() == this.min {
		for i := 0; i < len(this.stack)-1; i++ {
			if i == 0 || this.stack[i] < this.min {
				this.min = this.stack[i]
			}
		}
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
