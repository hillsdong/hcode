package leetcode

// 哈希
func twoSum(nums []int, target int) []int {
	l := len(nums)
	if l < 2 {
		return []int{}
	}

	numsHash := map[int]int{}
	for i := 0; i < l; i++ {
		if j, ok := numsHash[target-nums[i]]; ok {
			return []int{j, i}
		}
		numsHash[nums[i]] = i
	}
	return []int{}
}

// 暴力法
func twoSum2(nums []int, target int) []int {
	l := len(nums)
	if l < 2 {
		return []int{}
	}

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
