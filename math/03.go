package math

import (
	"fmt"
	"math"
)

//Q03 求一个大于零的数的平方根
func Q03(a float64) float64 {
	if a < 0 {
		return a
	}
	b := a
	for math.Abs(b*b-a) > 0.05 {
		b = (b + a/b) / 2
		fmt.Println(b)
	}
	return b
}
