package leetcode

import "fmt"

var rowNumHashs, colNumHashs, blkNumHashs []map[byte]bool

type sudokuPos struct{ row, col int }

var solvePos []sudokuPos

func solveSudoku(board [][]byte) {
	solveSudokuInit(board)
	_solveSudoku(board, 0)
}
func _solveSudoku(board [][]byte, posi int) {
	n := len(board)
	// termianl
	if posi == len(solvePos) {
		//printSoduku(board)
		return
	}

	row, col := solvePos[posi].row, solvePos[posi].col
	for i := 1; i <= n; i++ {
		ib := byte(48 + i)
		// cut
		if !isValidSudokuNum(row, col, ib) {
			continue
		}
		// program
		board[row][col] = ib
		// recursion
		_solveSudoku(board, posi+1)
		// recover
		rowNumHashs[row][ib] = false
		colNumHashs[col][ib] = false
		blkNumHashs[3*(row/3)+col/3][ib] = false
	}
}

func solveSudokuInit(board [][]byte) {
	n := len(board)
	rowNumHashs = make([]map[byte]bool, n)
	colNumHashs = make([]map[byte]bool, n)
	blkNumHashs = make([]map[byte]bool, n)
	for i := 0; i < n; i++ {
		rowNumHashs[i] = map[byte]bool{}
		colNumHashs[i] = map[byte]bool{}
		blkNumHashs[i] = map[byte]bool{}
	}
	solvePos = []sudokuPos{}
	var num byte

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			num = board[row][col]
			if num == '.' {
				solvePos = append(solvePos, sudokuPos{row, col})
				continue
			}
			rowNumHashs[row][num] = true
			colNumHashs[col][num] = true
			blkNumHashs[3*(row/3)+col/3][num] = true
		}
	}
	return
}

func isValidSudokuNum(row int, col int, num byte) bool {
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

	return true
}

func printSoduku(nums [][]byte) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			fmt.Printf("%q ", nums[i][j])
		}
		fmt.Println("")
	}
}
