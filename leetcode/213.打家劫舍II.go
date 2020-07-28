package leetcode

func robb(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp0, dp1 := 0, 0
	for i := 1; i < len(nums); i++ {
		dp0, dp1 = dp1, max(dp1, dp0+nums[i])
	}
	dp3, dp4 := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		dp3, dp4 = dp4, max(dp4, dp3+nums[i])
	}
	return max(dp1, dp4)
}
