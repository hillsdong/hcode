package leetcode

// TODO 分治

// 动态规划
func maxSubArray(nums []int) int {
	var sum int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > sum {
			sum = nums[i]
		}
	}
	return sum
}
