package leetcode

import (
	"container/list"
	"fmt"
)

type trap struct{}

// TODO 暴力

// TODO 哈希

// TODO 双指针

// 栈
func (t trap) Do(height []int) int {
	stack := list.New()
	stack.PushBack(0)

	var ret int
	for i := 1; i < len(height); i++ {
		if height[i] > height[i-1] {

			for stack.Len() != 0 {
				last := stack.Back()
				leftI := last.Value.(int)
				if height[leftI] >= height[i] {
					break
				}
				stack.Remove(last)
				if stack.Len() == 0 {
					break
				}
				ret += (min(height[i], height[stack.Back().Value.(int)]) - height[leftI]) * (i - 1 - stack.Back().Value.(int))
				fmt.Println(ret, leftI, i)
			}
		}
		stack.PushBack(i)
	}

	return ret
}
