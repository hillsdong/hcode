package find

import (
	"testing"
)

func TestFind(t *testing.T) {
	a := "o vosld! hello world!"
	b := "wor"
	t.Log(Find(a, b))
}

func BenchmarkFind(b *testing.B) {
	m := "o vosld! hello world!"
	p := "worl"
	for i := 0; i < b.N; i++ {
		Find(m, p)
	}
}

func TestFind2(t *testing.T) {
	a := "o vosld! hello world!"
	b := "wor"
	t.Log(Find2(a, b))
}

func BenchmarkFind2(b *testing.B) {
	m := "o vosld! hello world!"
	p := "wor"
	for i := 0; i < b.N; i++ {
		Find2(m, p)
	}
}
