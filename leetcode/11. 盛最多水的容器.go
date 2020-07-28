package leetcode

// 双指针
func maxArea(height []int) int {
	var a, maxA int
	i, j := 0, len(height)-1
	for i < j {
		a = (j - i) * min(height[i], height[j])

		if maxA < a {
			maxA = a
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxA
}

// 暴力 优化
func maxArea3(height []int) int {
	var a, h, maxA int
	for i := 0; i < len(height)-1; i++ {
		for j := len(height) - 1; j > i; j-- {
			if j < len(height)-1 && height[j] < height[j+1] {
				continue
			}
			h = height[i]
			if height[j] < h {
				h = height[j]
			}
			a = (j - i) * h
			if maxA < a {
				maxA = a
			}
		}
	}
	return maxA
}

// 暴力
func maxArea2(height []int) int {
	var a, h, maxA int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			h = height[i]
			if height[j] < h {
				h = height[j]
			}
			a = (j - i) * h
			if maxA < a {
				maxA = a
			}
		}
	}
	return maxA
}
