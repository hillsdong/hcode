package recursion

import "testing"

func TestFac(t *testing.T) {
	t.Log(Fac(3))
	t.Log(Fac(4))
	t.Log(Fac(10))
}

func TestFac2(t *testing.T) {
	t.Log(Fac2(3))
	t.Log(Fac2(4))
	t.Log(Fac2(10))
}
