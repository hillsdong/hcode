package find

import (
	"testing"
)

func TestFind(t *testing.T) {
	a := "orlorlworlworl"
	b := "worl"
	t.Log(Find(a, b))
}

func TestGenerateBC(t *testing.T) {
	b := "orlorlworlworl"
	t.Log(generateBC(b))
}

func TestGenerateGS(t *testing.T) {
	b := "orlorlworlworl"
	t.Log(generateGS(b))
}

func BenchmarkFind(b *testing.B) {
	m := "o vosld! hello world!"
	p := "worl"
	for i := 0; i < b.N; i++ {
		Find(m, p)
	}
}
