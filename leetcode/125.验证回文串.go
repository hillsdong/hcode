package leetcode

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {

		for !isAN(s[i]) && i < j {
			i++
		}
		for !isAN(s[j]) && i < j {
			j--
		}

		if i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}

func isAN(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
