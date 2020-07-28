package leetcode

import "testing"

func TestTrap(t *testing.T) {
	//t.Log(new(trap).Do([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	//t.Log(new(trap).Do([]int{2, 1, 0, 2}))                         //3
	//t.Log(new(trap).Do([]int{4, 2, 3}))                            // 1
	t.Log(new(trap).Do([]int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6})) //23
}
