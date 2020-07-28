package leetcode

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	t.Log(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) //10
	t.Log(largestRectangleArea([]int{1}))                //1
	t.Log(largestRectangleArea([]int{1, 1}))             //2
	
}
