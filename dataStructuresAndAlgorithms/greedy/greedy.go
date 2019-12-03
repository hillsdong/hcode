package greedy

import "fmt"

var total = 0

// NQueens 求解N皇后问题
func NQueens(n int) {
	queens := make([]int, n)
	nQueens(queens, 0)
}

func nQueens(queens []int, i int) {
	n := len(queens)
	if i == n {
		printQueens(queens)
	}
	for j := 0; j < n; j++ {
		if isOk(queens, i, j) {
			queens[i] = j
			nQueens(queens, i+1)
		}
	}
}

func isOk(queens []int, i int, j int) bool {
	for row := i - 1; row > -1; row-- {
		if queens[row] == j || queens[row] == j-(i-row) || queens[row] == j+(i-row) {
			return false
		}
	}
	return true
}

func printQueens(queens []int) {
	n := len(queens)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if queens[i] == j {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Print("\n")
	}
	total++
	fmt.Printf("%d", total)
	for i := 0; i < n; i++ {
		fmt.Print("- ")
	}
	fmt.Print("\n")
}
