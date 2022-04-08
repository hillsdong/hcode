package leetcode

func longestPalindrome(s string) string {
	n := len(s)
	ret := ""
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if i == j {
				dp[i][j] = true
			} else if i+1 == j {
				if s[i] == s[j] {
					dp[i][j] = true
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && l+1 > len(ret) {
				ret = s[i : j+1]
			}
		}
	}
	return ret
}

func longestPalindrome2(s string) string {
	n := len(s)
	ret := ""
	for i := 0; i < n; i++ {
		j := 0
		for ; i+j < n && i-j >= 0; j++ {
			if s[i+j] != s[i-j] {
				break
			}
		}
		if 2*j-1 > len(ret) {
			ret = s[i-j+1 : i+j]
		}
		j = 0
		for ; i+j+1 < n && i-j >= 0; j++ {
			if s[i+j+1] != s[i-j] {
				break
			}
		}
		if 2*j > len(ret) {
			ret = s[i-j+1 : i+j+1]
		}

	}
	return ret
}
