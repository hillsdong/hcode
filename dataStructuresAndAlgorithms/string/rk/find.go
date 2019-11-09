package find

// Find return the first index that b in a, realized the RK string search algo
func Find(a, b string) int {
	n, m := len(a), len(b)
	if n == 0 || m == 0 || m > n {
		return -1
	}

	var ah, bh int
	for j := 0; j < m; j++ {
		bh = bh + int(b[j])
	}

	for i := 0; i < n-m+1; i++ {

		if i == 0 {
			for j := 0; j < m; j++ {
				ah = ah + int(a[j])
			}
		} else {
			ah = ah + int(a[i+m-1]) - int(a[i-1])
		}

		//fmt.Printf("ah: %d, bh: %d, a: %s, b: %s \n", ah, bh, a[i:i+m], b)
		if ah == bh && a[i:i+m] == b {
			return i
		}
	}
	return -1
}

// Find2 return the first index that b in a, realized the RK string search algo
func Find2(a, b string) int {
	n, m := len(a), len(b)
	if n == 0 || m == 0 || m > n {
		return -1
	}

	var ah, bh int
	bh = hash(b, 0, 0, m)

	for i := 0; i < n-m+1; i++ {
		ah = hash(a, ah, i, m)
		//fmt.Printf("ah: %d, bh: %d, a: %s, b: %s \n", ah, bh, a[i:i+m], b)
		if ah == bh && a[i:i+m] == b {
			return i
		}
	}
	return -1
}

func hash(a string, ah, i, m int) int {
	if i == 0 {
		for j := 0; j < m; j++ {
			ah = ah + int(a[j])
		}
	} else {
		ah = ah + int(a[i+m-1]) - int(a[i-1])
	}
	return ah
}
