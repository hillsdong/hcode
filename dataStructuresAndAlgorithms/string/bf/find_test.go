package find

import "testing"

func TestFind(t *testing.T) {
	a := "o vosld! hello world!"
	b := "worl"
	t.Log(Find(a, b))
}

func BenchmarkFind(b *testing.B) {
	m := "o vosld! hello world!"
	p := "worl"
	for i := 0; i < b.N; i++ {
		Find(m, p)
	}
}
