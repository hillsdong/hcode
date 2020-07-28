package leetcode

// 一维
func minimumTotal(triangle [][]int) int {
	cc := triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			cc[j] = triangle[i][j] + min(cc[j], cc[j+1])
		}
	}

	return cc[0]
}

// 二维
func minimumTotal2(triangle [][]int) int {
	cc := triangle
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			cc[i][j] = triangle[i][j] + min(cc[i+1][j], cc[i+1][j+1])
		}
	}

	return cc[0][0]
}
