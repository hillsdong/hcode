package img

import "testing"

func TestCut(t *testing.T) {
	i := New("s.png")
	t.Log(i)

	i.Cut(-380, -350, 380, 350)

	t.Log(i)
	i.Save("s100.png")
}

func TestCircle(t *testing.T) {
	i := New("s.png")
	t.Log(i)

	i.Circle(-190, -175, 175)

	t.Log(i)
	i.Save("s200.png")
}
