package main

import "testing"

//TestMode for test mode
func TestMode(t *testing.T) {
	cases := []struct {
		in, want []float64
	}{
		{[]float64{1}, []float64{}},
		{[]float64{1, 3, 4}, []float64{}},
		{[]float64{1, 2, 2, 3, 3, 4}, []float64{2, 3}},
	}
	for _, c := range cases {
		got := mode(c.in)
		if !Float64SliceEqual(got, c.want) {
			t.Errorf("Reverse(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

//Float64SliceEqual for compare float64 slice
func Float64SliceEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func BenchmarkMode(b *testing.B) {
	nums := []float64{3, 4, 4.5, 5, 5, 6.2, 7.1, 7.1, 8.5, 9}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mode(nums)
	}
}

func BenchmarkMode2(b *testing.B) {
	nums := []float64{3, 4, 4.5, 5, 5, 6.2, 7.1, 7.1, 8.5, 9}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mode2(nums)
	}
}

func TestQuadratic(t *testing.T) {
	x1, x2 := quadratic(1, 2, 2)
	t.Errorf("%v,%v", x1, x2)
}
