package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("babad"))
	t.Log(longestPalindrome("bababd"))
	t.Log(longestPalindrome("cbbd"))
	t.Log(longestPalindrome("abcda"))
}
