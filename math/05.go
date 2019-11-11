package math

import "fmt"

//Q05 求一个整数的所有乘积因子
func Q05(a int, res []int) {
	if a == 1 {
		fmt.Println(res)
	}
	for i := 2; i < a+1; i++ {
		if a%i == 0 {
			res := res
			res = append(res, i)
			Q05(a/i, res)
		}
	}
}
