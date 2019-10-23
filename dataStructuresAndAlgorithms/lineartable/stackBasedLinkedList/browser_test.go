package stack

import "testing"

func TestBrowser(t *testing.T) {
	b := NewBrowser()
	b.Push("www.qq.com")
	t.Log((b))
	b.Push("www.baidu.com")
	b.Push("www.sina.com")
	t.Log((b))

	b.Back()
	t.Log((b))
	b.Forward()

	t.Log((b))

}
