package recursion

import "testing"

func TestFib(t *testing.T) {
	t.Log(Fib(3))
	t.Log(Fib(4))
	t.Log(Fib(30))
}

func TestFib2(t *testing.T) {
	t.Log(Fib2(3))
	t.Log(Fib2(4))
	t.Log(Fib2(30))
}

func TestFib3(t *testing.T) {
	t.Log(Fib3(3))
	t.Log(Fib3(4))
	t.Log(Fib3(30))
}
