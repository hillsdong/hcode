package sort

import "fmt"

// MaxK return the kth max number in a numbers array
func MaxK(nums []int, k int) int{
	return _maxK(nums, 0, len(nums)-1, k-1)
}

func _maxK(nums []int, p int, r int, k int) int {
	if k > r || k < p{
		return -1
	}
	//分区
	q := p
	for i := p; i <= r; i++ {
		if nums[i] > nums[r] {
			nums[i], nums[q] = nums[q], nums[i]
			q++
		}
	}
	nums[q], nums[r] = nums[r], nums[q]
	//分治
	if q == k {
		return nums[k]
	} else if q > k {
		return _maxK(nums, p, q-1, k)
	}else{
		return _maxK(nums, q+1, r, k)
	}
}