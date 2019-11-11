package math
	
import "testing"
import "fmt"
import "math"

func Test_Q03_1(t *testing.T) {
	var a float64 = 64
	b := Q03(a)
	if math.Abs(b*b - a) < 0.05 {
		t.Log("Q03_1 pass")
	} else {
		t.Error("Q03_1 err "+ fmt.Sprint(b) )
	}
}

func Test_Q03_2(t *testing.T) {
	a := 0.5
	b := Q03(a)
	if math.Abs(b*b - a) < 0.05 {
		t.Log("Q03_2 pass")
	} else {
		t.Error("Q03_2 err "+ fmt.Sprint(b) )
	}
}

func Test_Q03_3(t *testing.T) {
	a := 1.5
	b := Q03(a)
	if math.Abs(b*b - a) < 0.05 {
		t.Log("Q03_3 pass")
	} else {
		t.Error("Q03_3 err "+ fmt.Sprint(b) )
	}
}