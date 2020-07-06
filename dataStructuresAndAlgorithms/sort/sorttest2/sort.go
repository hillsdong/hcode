package sorttest2

// 2min
func Bubble(nums []int) []int {
	len := len(nums)
	if len <= 1 {
		return nums
	}

	for i := len; i > 0; i-- {
		for j := 1; j < i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}

// 3min
func Insertion(nums []int) []int {
	len := len(nums)
	if len <= 1 {
		return nums
	}

	for i := 1; i < len; i++ {
		v := nums[i]
		j := i
		for ; j > 0; j-- {
			if nums[j-1] > v {
				nums[j] = nums[j-1]
			} else {
				break
			}
		}
		nums[j] = v
	}
	return nums
}

// 4min
func Selection(nums []int) []int {
	len := len(nums)
	if len <= 1 {
		return nums
	}

	for i := 0; i < len; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

// 10min
func Merge(nums []int) []int {
	_Merge(nums, 0, len(nums)-1)
	return nums
}

func _Merge(nums []int, p int, r int) {
	len := r - p + 1
	if len <= 1 {
		return
	}

	q := p + (r-p)/2

	_Merge(nums, p, q)
	_Merge(nums, q+1, r)

	_nums, k, i, j := make([]int, len), 0, p, q+1

	for ; i <= q && j <= r; k++ {
		if nums[i] < nums[j] {
			_nums[k] = nums[i]
			i++
		} else {
			_nums[k] = nums[j]
			j++
		}
	}

	for ; i <= q; k++ {
		_nums[k] = nums[i]
		i++
	}

	for ; j <= r; k++ {
		_nums[k] = nums[j]
		j++
	}

	for i := 0; i < k; i++ {
		nums[p+i] = _nums[i]
	}
	return
}

// 14min
func Quick(nums []int) []int {
	_Quick(nums, 0, len(nums)-1)
	return nums
}

func _Quick(nums []int, p int, r int) {
	len := r - p + 1
	if len <= 1 {
		return
	}

	q := p
	for i := p; i < r; i++ {
		if nums[i] < nums[r] {
			nums[i], nums[q] = nums[q], nums[i]
			q++
		}
	}
	nums[q], nums[r] = nums[r], nums[q]

	_Quick(nums, p, q-1)
	_Quick(nums, q+1, r)
	return
}
