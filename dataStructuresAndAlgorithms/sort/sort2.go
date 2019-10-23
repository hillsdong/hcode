package sort

func merge(nums []int) []int {
	_merge(nums, 0, len(nums)-1)
	return nums
}
func _merge(nums []int, p int, r int) {
	n := r - p + 1
	if n <= 1 {
		return
	}
	//分治
	q := p + (r-p)/2
	_merge(nums, p, q)
	_merge(nums, q+1, r)
	//合并
	_nums, i, j, k := make([]int, n), p, q+1, 0
	for ; k < n; k++ {
		if i > q {
			_nums[k] = nums[j]
			j++
		} else if j > r {
			_nums[k] = nums[i]
			i++
		} else if nums[i] < nums[j] {
			_nums[k] = nums[i]
			i++
		} else {
			_nums[k] = nums[j]
			j++
		}
	}
	for i := 0; i < n; i++ {
		nums[p+i] = _nums[i]
	}
}

func quik(nums []int) []int {
	_quik(nums, 0, len(nums)-1)
	return nums
}
func _quik(nums []int, p int, r int) {
	if p >= r {
		return
	}
	//分区
	q := p
	for i := p; i <= r; i++ {
		if nums[i] < nums[r] {
			nums[i], nums[q] = nums[q], nums[i]
			q++
		}
	}
	nums[q], nums[r] = nums[r], nums[q]
	//分治
	_quik(nums, p, q-1)
	_quik(nums, q+1, r)
}
