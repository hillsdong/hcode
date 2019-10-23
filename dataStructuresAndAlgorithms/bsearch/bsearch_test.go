package bsearch

import "testing"

func nums() []int {
	return []int{2, 3, 4, 8, 8, 8, 8, 9, 10}
}

func TestB1(t *testing.T) {
	t.Error(b1(nums(), 10))
}

func TestB2(t *testing.T) {
	t.Error(b2(nums(), 7))
}

func TestB3(t *testing.T) {
	t.Error(b3(nums(), 8))
}

func TestB4(t *testing.T) {
	t.Error(b4(nums(), 8))
}

func TestB5(t *testing.T) {
	t.Error(b5(nums(), 8))
}

func TestB6(t *testing.T) {
	t.Error(b6(nums(), 8))
}
