package find

import "fmt"

// Find return the first index that b in a, realized the BM string search algo
func Find(a, b string) int {
	n, m := len(a), len(b)
	if n == 0 || m == 0 || m > n {
		return -1
	}

	bc := generateBC(b)
	suffix, prefix := generateGS(b)

	var step, gsStep int
	for i := 0; i < n-m+1; i += step {
		j := m - 1
		//find
		for ; a[i+j] == b[j]; j-- {
			if j == 0 {
				return i
			}
		}
		//bc step
		step = j - (bc[a[i+j]] - 1)
		bcStep := step // for log
		gsStep = 0     //for log
		//gs step
		if j < m-1 {
			j = j + 1 //good suffix's first index
			gsStep = j - (suffix[j] - 1)
			if suffix[j] == 0 {
				for j = j + 1; j < m-1; j++ {
					if prefix[j] {
						gsStep = j
					}
				}
			}
			if gsStep > step {
				step = gsStep
			}
		}

		if step < 1 {
			step = 1
		}

		fmt.Printf("bc: %s, step: %d, bcStep: %d,  gsStep: %d, a: %s\n", string(a[i+j]), step, bcStep, gsStep, a[0:i+m])
	}
	return -1
}

func generateBC(b string) map[byte]int {
	bc := make(map[byte]int)
	for i := 0; i < len(b); i++ {
		bc[b[i]] = i + 1
	}
	return bc
}

// generateGS generate suffix and prefix map
// suffix is a map from the first index of good suffix to the first index of last same string
// prefix is a map from the first index of good suffix to a bool that is same with the prefix of string
func generateGS(b string) ([]int, []bool) {
	m := len(b)
	suffix := make([]int, m)
	prefix := make([]bool, m)

	for i := 0; i < m-1; i++ {
		j := i
		k := m - 1
		for j >= 0 && b[j] == b[k] {
			suffix[k] = j
			if j == 0 {
				prefix[k] = true
			}
			j--
			k--
		}
	}

	return suffix, prefix
}
