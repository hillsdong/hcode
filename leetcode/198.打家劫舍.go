package leetcode

func rob(nums []int) int {
	dp0, dp1 := 0, 0
	for i := 0; i < len(nums); i++ {
		dp0, dp1 = dp1, max(dp1, dp0+nums[i])
	}
	return dp1
}

func rob4(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp0, dp1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp0, dp1 = dp1, max(dp1, dp0+nums[i])
	}
	return dp1
}

func rob3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	ret := max(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		ret = max(ret, dp[i])
	}
	return ret
}

var cc map[bool]map[int]int

func robinit() {
	cc = map[bool]map[int]int{
		true:  map[int]int{},
		false: map[int]int{},
	}
}
func rob1(nums []int) int {
	robinit()
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	return max(_rob(nums, 0, true), _rob(nums, 0, false))
}

func _rob(nums []int, i int, r bool) int {
	if v, ok := cc[r][i]; ok {
		return v
	}

	if i >= len(nums) {
		cc[r][i] = 0
		return cc[r][i]
	}
	if i == len(nums)-1 {
		if r {
			cc[r][i] = nums[i]
			return cc[r][i]
		}
		cc[r][i] = 0
		return cc[r][i]
	}
	if r {
		cc[r][i] = nums[i] + max(_rob(nums, i+2, true), _rob(nums, i+2, false))
		return cc[r][i]
	}
	cc[r][i] = max(_rob(nums, i+1, true), _rob(nums, i+1, false))
	return cc[r][i]
}
