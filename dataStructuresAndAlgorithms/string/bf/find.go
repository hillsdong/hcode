package find

// Find return the first index that b in a, realized the BF string search algo
func Find(a, b string) int {
	n, m := len(a), len(b)
	if n == 0 || m == 0 || m > n {
		return -1
	}
	for i := 0; i < n-m+1; i++ {
		if a[i:i+m] == b {
			return i
		}
	}
	return -1
}
