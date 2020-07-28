package leetcode

// TODO 动态规划

// 贪心
func maxProfit(prices []int) int {
	var minPirce int = prices[0]
	var max int
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPirce {
			minPirce = prices[i]
		}
		if prices[i]-minPirce > max {
			max = prices[i] - minPirce
		}
	}
	return max
}

// 暴力
func maxProfit2(prices []int) int {
	var max int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}
