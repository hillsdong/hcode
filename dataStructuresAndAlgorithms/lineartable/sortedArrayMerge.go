package lineartable

// SortedArrayMerge 合并两个有序数组
func SortedArrayMerge(a []int64, b []int64) []int64 {
	if len(a) < 1 {
		return b
	}
	if len(b) < 1 {
		return a
	}
	c, i, j := make([]int64, len(a)+len(b)), 0, 0
	for k := 0; k < len(a)+len(b); k++ {
		if i >= len(a) {
			c[k] = b[j]
			j++
		} else if j >= len(b) {
			c[k] = a[i]
			i++
		} else if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
	}
	return c
}

// SortedArrayMerge2 合并两个有序数组2
func SortedArrayMerge2(a []int64, b []int64) []int64 {
	if len(a) < 1 {
		return b
	}
	if len(b) < 1 {
		return a
	}
	c, i, j, k := make([]int64, len(a)+len(b)), 0, 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[k] = a[i]
			k++
			i++
		} else {
			c[k] = b[j]
			k++
			j++
		}
	}

	for i < len(a) {
		c[k] = a[i]
		k++
		i++
	}
	for j < len(b) {
		c[k] = b[j]
		k++
		j++
	}
	return c
}
