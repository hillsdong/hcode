package img

import "testing"

func TestLoad(t *testing.T) {
	i := New("s.png")
	t.Log(i)

	i.Cut(-380, -350, 380, 350)

	t.Log(i)
	i.Save("s100.png")
}
