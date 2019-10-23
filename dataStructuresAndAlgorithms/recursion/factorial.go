package recursion

// Fac return n's factorial by recursion
func Fac(n int) int {
	if n < 2 {
		return n
	}
	return n * Fac(n-1)
}

// Fac2 return n's factorial by foreach
func Fac2(n int) (ret int) {
	ret = 1
	for i := 2; i <= n; i++ {
		ret = ret * i
	}
	return
}
