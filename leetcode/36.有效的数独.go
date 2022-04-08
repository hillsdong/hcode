package leetcode

func isValidSudoku(nums [][]byte) bool {
	n := len(nums)
	rowNumHashs := make([]map[byte]bool, n)
	colNumHashs := make([]map[byte]bool, n)
	blkNumHashs := make([]map[byte]bool, n)
	for i := 0; i < n; i++ {
		rowNumHashs[i] = map[byte]bool{}
		colNumHashs[i] = map[byte]bool{}
		blkNumHashs[i] = map[byte]bool{}
	}
	var num byte

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			num = nums[row][col]
			if num == '.' {
				continue
			}
			if rowNumHashs[row][num] {
				return false
			}
			if colNumHashs[col][num] {
				return false
			}
			if blkNumHashs[3*(row/3)+col/3][num] {
				return false
			}
			rowNumHashs[row][num] = true
			colNumHashs[col][num] = true
			blkNumHashs[3*(row/3)+col/3][num] = true
		}
	}
	return true
}

func isValidSudoku2(nums [][]byte) bool {
	var numHash map[byte]bool
	var num byte
	n := len(nums)

	for row := 0; row < n; row++ {
		numHash = map[byte]bool{}
		for col := 0; col < n; col++ {
			num = nums[row][col]
			if num != '.' && numHash[num] {
				return false
			}
			numHash[num] = true
		}
	}
	for col := 0; col < n; col++ {
		numHash = map[byte]bool{}
		for row := 0; row < n; row++ {
			num = nums[row][col]
			if num != '.' && numHash[num] {
				return false
			}
			numHash[num] = true
		}
	}
	for block := 0; block < (n/3)*(n/3); block++ {
		numHash = map[byte]bool{}
		for row := 0; row < n/3; row++ {
			for col := 0; col < n/3; col++ {
				num = nums[3*(block/3)+row][3*(block%3)+col]
				if num != '.' && numHash[num] {
					return false
				}
				numHash[num] = true
			}
		}
	}
	return true
}
