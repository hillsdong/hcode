package leetcode

// 递归 缓存
var fiboCache map[int]int

func fiboInit() {
	fiboCache = map[int]int{}
}
func fibonacci(d, n int) int {
	if n < 2 {
		return n
	}

	if v, exist := fiboCache[n]; exist {
		return v
	}
	fiboCache[n] = fibonacci(d+1, n-1) + fibonacci(d+1, n-2)
	return fiboCache[n]
}

// 递归
func fibonacci3(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci3(n-1) + fibonacci3(n-2)
}

// 递推
func fibonacci2(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
