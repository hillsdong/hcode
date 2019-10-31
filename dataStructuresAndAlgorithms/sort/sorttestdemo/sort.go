package sorttestdemo

func Bubble(nums []int) []int {
	len := len(nums)
	if len <= 1 {
		return nums
	}

	return nums
}

func Insertion(nums []int) []int {
	len := len(nums)
	if len <= 1 {
		return nums
	}

	return nums
}

func Selection(nums []int) []int {
	len := len(nums)
	if len <= 1 {
		return nums
	}

	return nums
}

func Merge(nums []int) []int {
	_Merge(nums, 0, len(nums)-1)
	return nums
}

func _Merge(nums []int, p int, r int) {
	len := r - p + 1
	if len <= 1 {
		return
	}
	return
}

func Quick(nums []int) []int {
	_Quick(nums, 0, len(nums)-1)
	return nums
}

func _Quick(nums []int, p int, r int) {
	len := r - p + 1
	if len <= 1 {
		return
	}
	return
}
