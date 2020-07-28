package leetcode

import "fmt"

// 最清晰简洁
func moveZeroes(nums []int) {
	zi := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zi] = nums[i]
			zi++
		}
	}
	for i := zi; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 大学 更清晰简洁
func moveZeroes4(nums []int) {
	zi, nzi := -1, -1

	for i := 0; i < len(nums); i++ {
		fmt.Println(zi, nzi, i, nums)

		if zi == -1 && nums[i] == 0 {
			zi = i
		}

		if zi != -1 && nzi == -1 && nums[i] != 0 {
			nzi = i
		}

		if zi != -1 && nzi != -1 && nums[i] != 0 {
			nums[i], nums[zi] = nums[zi], nums[i]
			zi++
		}
	}
}

// 大学 更清晰
func moveZeroes3(nums []int) {
	zi, nzi := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zi = i
			break
		}
	}
	if zi == -1 {
		return
	}
	for i := zi + 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nzi = i
			break
		}
	}
	if nzi == -1 {
		return
	}
	for i := nzi; i < len(nums); i++ {
		fmt.Println(zi, i, nums)

		if nums[i] != 0 {
			nums[i], nums[zi] = nums[zi], nums[i]
			zi++
		}
	}
}

// 大学 稍改进 最牛逼
func moveZeroes2(nums []int) {
	for i, zi := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zi] = nums[zi], nums[i]
			zi++
		}
	}
}
