package math

import "testing"

func Test_Q01_1(t *testing.T) {
	binaryStr := Q01(25)
	if binaryStr == "00000000000000000000000000011001" {
		t.Log("Q01_1 pass")
	} else {
		t.Error("Q01_1 err " + binaryStr)
	}
}

func Test_Q01_2(t *testing.T) {
	binaryStr := Q01(-25)
	if binaryStr == "11111111111111111111111111100111" {
		t.Log("Q01_2 pass")
	} else {
		t.Error("Q01_2 err " + binaryStr)
	}
}
