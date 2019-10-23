package bsearch

func b1(nums []int, num int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if nums[mid] == num {
			return mid
		} else if nums[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func b2(nums []int, num int) int {
	return _b2(nums, num, 0, len(nums)-1)
}
func _b2(nums []int, num int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := low + ((high - low) >> 1)
	if nums[mid] == num {
		return mid
	} else if nums[mid] < num {
		return _b2(nums, num, mid+1, high)
	} else {
		return _b2(nums, num, low, mid-1)
	}
}

//第一个等于
func b3(nums []int, num int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if nums[mid] == num {
			if mid == 0 || nums[mid-1] < num {
				return mid
			}
			high = mid - 1
		} else if nums[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

//最后一个等于
func b4(nums []int, num int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if nums[mid] == num {
			if mid == len(nums)-1 || nums[mid+1] > num {
				return mid
			}
			low = mid - 1
		} else if nums[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

//第一个大于等于
func b5(nums []int, num int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if nums[mid] >= num {
			if mid == 0 || nums[mid-1] < num {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//最后一个小于等于
func b6(nums []int, num int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + ((high - low) >> 1)
		if nums[mid] <= num {
			if mid == len(nums)-1 || nums[mid+1] > num {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
