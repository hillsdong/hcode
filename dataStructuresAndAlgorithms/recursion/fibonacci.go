package recursion

// FidHash stores fid value
var FidHash map[int]int

func init() {
	FidHash = make(map[int]int)
}

// Fib return the nth fibonacci number by recursion
func Fib(n int) int {
	if n < 3 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// Fib2 return the nth fibonacci number by foreach
func Fib2(n int) int {
	if n < 3 {
		return n
	}
	tmp := 1
	sum := 2
	pre := tmp

	for i := 2; i < n; i++ {
		tmp = sum
		sum = pre + sum
		pre = tmp
	}
	return sum
}

// Fib3 return the nth fibonacci number by recursion
func Fib3(n int) int {
	if v, ok := FidHash[n]; ok {
		return v
	}

	if n < 3 {
		FidHash[n] = n
	} else {
		FidHash[n] = Fib(n-1) + Fib(n-2)
	}

	return FidHash[n]
}
