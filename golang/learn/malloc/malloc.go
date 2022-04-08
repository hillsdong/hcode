package main

import (
	"math/rand"
	"runtime"
	"time"
)

const CAP = 10000000

type AdFeature struct {
	a int
	b int
	c int
	d int
	e int
	f int
}

var FeatureMap interface{}

func value() interface{} {
	m := make(map[int]AdFeature, CAP)
	for i := 0; i < CAP; i++ {
		adFea := AdFeature{}
		adFea.a = rand.Intn(1000000)
		adFea.b = rand.Intn(1000000)
		adFea.c = rand.Intn(1000000)
		adFea.d = rand.Intn(1000000)
		adFea.e = rand.Intn(1000000)
		adFea.f = rand.Intn(1000000)
		m[i] = adFea
	}
	return m
}

func pointer() interface{} {
	m := make(map[int]*AdFeature, CAP)
	for i := 0; i < CAP; i++ {
		adFea := &AdFeature{}
		adFea.a = rand.Intn(1000000)
		adFea.b = rand.Intn(1000000)
		adFea.c = rand.Intn(1000000)
		adFea.d = rand.Intn(1000000)
		adFea.e = rand.Intn(1000000)
		adFea.f = rand.Intn(1000000)
		m[i] = adFea
	}
	return m
}

func main() {
	FeatureMap = pointer() //value()
	for i := 0; i < 20; i++ {
		runtime.GC()
		time.Sleep(time.Second)
	}
}

//GODEBUG="gctrace=1" go run malloc.go
