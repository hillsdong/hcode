package leetcode

// 栈+哈希 基于高度优化 再优化
func largestRectangleArea(heights []int) int {
	var maxArea, area int
	s := stack{}
	left, right := map[int]int{}, map[int]int{}
	for i := 0; i < len(heights); i++ {
		right[i] = len(heights)
	}
	for i := 0; i < len(heights); i++ {
		for !s.Empty() && heights[s.Top()] >= heights[i] {
			right[s.Top()] = i
			s.Pop()
		}
		if len(s) == 0 {
			left[i] = -1
		} else {
			left[i] = s.Top()
		}
		s.Push(i)
	}
	for i := 0; i < len(heights); i++ {
		area = (right[i] - left[i] - 1) * heights[i]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// 栈+哈希 基于高度优化
func largestRectangleArea4(heights []int) int {
	var maxArea, area int
	s := stack{}
	left, right := map[int]int{}, map[int]int{}
	for i := 0; i < len(heights); i++ {
		for !s.Empty() && heights[s.Top()] >= heights[i] {
			s.Pop()
		}
		if len(s) == 0 {
			left[i] = -1
		} else {
			left[i] = s.Top()
		}
		s.Push(i)
	}
	s = stack{}
	for i := len(heights) - 1; i >= 0; i-- {
		for !s.Empty() && heights[s.Top()] >= heights[i] {
			s.Pop()
		}
		if len(s) == 0 {
			right[i] = len(heights)
		} else {
			right[i] = s.Top()
		}
		s.Push(i)
	}
	for i := 0; i < len(heights); i++ {
		area = (right[i] - left[i] - 1) * heights[i]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// 暴力 基于高度
func largestRectangleArea3(heights []int) int {
	var maxArea, area, left, right int
	for i := 0; i < len(heights); i++ {
		left, right = i-1, i+1
		for l := i - 1; l >= -1; l-- {
			if l == -1 || heights[l] < heights[i] {
				left = l
				break
			}
		}
		for r := i + 1; r <= len(heights); r++ {
			if r == len(heights) || heights[r] < heights[i] {
				right = r
				break
			}
		}
		area = heights[i] * (right - left - 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// 暴力 基于宽度
func largestRectangleArea2(heights []int) int {
	var maxArea, area int
	for i := 0; i < len(heights); i++ {
		minHeight := heights[i]
		for j := i; j < len(heights); j++ {
			if heights[j] < minHeight {
				minHeight = heights[j]
			}
			area = (j - i + 1) * minHeight
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
