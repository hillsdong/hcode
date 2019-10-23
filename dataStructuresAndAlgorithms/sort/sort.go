package sort

func bubble(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	var f bool
	for i := 0; i < n; i++ {
		f = false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				f = true
			}
		}
		if !f {
			break
		}
	}
	return nums
}

func insertion(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	var v, j int
	for i := 0; i < n; i++ {
		v = nums[i]
		j = i
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

func selection(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	var min int
	for i := 0; i < n; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
	return nums
}
