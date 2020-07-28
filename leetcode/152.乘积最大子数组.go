package leetcode

// DP3
func maxProduct(nums []int) int {
	maxp, minp, ret := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxp, minp = minp, maxp
		}
		maxp = max(nums[i], maxp*nums[i])
		minp = min(nums[i], minp*nums[i])
		ret = max(maxp, ret)
	}
	return ret
}

// DP2
func maxProduct4(nums []int) int {
	maxp, minp, ret := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxp, minp = max(nums[i], minp*nums[i]), min(nums[i], maxp*nums[i])
		} else {
			maxp, minp = max(nums[i], maxp*nums[i]), min(nums[i], minp*nums[i])
		}
		ret = max(maxp, ret)
	}
	return ret
}

// DP
func maxProduct3(nums []int) int {
	maxp, minp, ret := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		_maxp := maxp
		maxp = max(nums[i], max(_maxp*nums[i], minp*nums[i]))
		minp = min(nums[i], min(_maxp*nums[i], minp*nums[i]))
		ret = max(maxp, ret)
	}
	return ret
}

// 暴力
func maxProduct2(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		ip := nums[i]
		if ip > max {
			max = ip
		}
		for j := i + 1; j < len(nums); j++ {
			ip *= nums[j]
			if ip > max {
				max = ip
			}
		}
	}
	return max
}
