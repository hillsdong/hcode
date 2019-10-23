package stack

import "fmt"

type Browser struct {
	backStack    *Stack
	forwardStack *Stack
}

func NewBrowser() *Browser {
	return &Browser{New(), New()}
}

func (b *Browser) Push(addr string) {
	b.backStack.Push(addr)
}

func (b *Browser) Forward() {
	if b.forwardStack.Len() == 0 {
		return
	}
	v, _ := b.forwardStack.Pop()
	b.backStack.Push(v)
}

func (b *Browser) Back() {
	if b.backStack.Len() == 0 {
		return
	}
	v, _ := b.backStack.Pop()
	b.forwardStack.Push(v)
}

func (b *Browser) String() string {
	return "\n" + fmt.Sprint(b.backStack) + "\n" + fmt.Sprint(b.forwardStack)
}
