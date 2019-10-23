package sort

import "testing"

func nums() []int {
	return []int{3, 2, 4, 1, 7, 5, 6}
	//return []int{6, 7, 8, 9, 1, 2, 3, 4}
}
func TestBubble(t *testing.T) {
	t.Log(bubble(nums()))
}

func TestInertion(t *testing.T) {
	t.Log(insertion(nums()))
}

func TestSelection(t *testing.T) {
	t.Log(selection(nums()))
}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubble(nums())
	}
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertion(nums())
	}
}
func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selection(nums())
	}
}
