package leetcode

import "testing"

func TestLJJS(t *testing.T) {
	lj := make([][]bool, 8)
	for i := 0; i < 8; i++ {
		lj[i] = make([]bool, 8)
	}
	lj[1][2] = true
	lj[1][6] = true
	lj[2][4] = true
	lj[3][0] = true
	lj[3][2] = true
	lj[3][5] = true
	lj[4][2] = true
	lj[5][3] = true
	lj[5][4] = true
	lj[5][6] = true
	lj[6][1] = true
	lj[6][5] = true
	t.Log(ljjs(lj))
}
