package leetcode

import (
	"container/list"
	"strings"
)

// 栈
func isValid2(s string) bool {
	h := map[byte]byte{'{': '}', '(': ')', '[': ']'}
	stack := list.New()
	for _, v := range []byte(s) {
		if _, exist := h[v]; exist {
			stack.PushBack(v)
		} else {
			last := stack.Back()
			if last == nil || h[last.Value.(byte)] != v {
				return false
			}
			stack.Remove(last)
		}
	}
	return stack.Len() == 0
}

// 替换空
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	r := strings.NewReplacer("{}", "", "[]", "", "()", "")
	half := len(s) / 2
	for i := 0; i < half; i++ {
		//s = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "{}", ""), "()", ""), "[]", "")
		if s == "" {
			return true
		}
		s = r.Replace(s)
	}
	return s == ""
}
