package triangle

import "testing"

func TestTri(t *testing.T) {
	pointsV := [][]int{
		[]int{5},
		[]int{7, 8},
		[]int{2, 3, 4},
		[]int{4, 9, 6, 1},
		[]int{2, 7, 9, 4, 5},
	}
	b := New(pointsV)
	t.Log(b.MinValue())
	t.Log(b.status)
}
