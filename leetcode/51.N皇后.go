package leetcode

var _solveNQueensRet [][]string

func solveNQueens(n int) [][]string {
	// 初始化结果
	_solveNQueensRet = [][]string{}
	// 初始化棋盘
	qs := make([]int, n)
	// 深度优先搜索
	_solveNQueens(qs, 0)
	// 返回
	return _solveNQueensRet
}

// 递归函数
func _solveNQueens(qs []int, row int) {
	n := len(qs)
	// 终止条件
	if row == n {
		_solveNQueensGeneralRet(qs)
		return
	}
	for column := 0; column < n; column++ {
		// 剪枝
		if !_solveNQueensIsVal(qs, row, column) {
			continue
		}
		// 保留结果
		qs[row] = column
		// 下一层
		_solveNQueens(qs, row+1)
	}
	return
}

// 检验函数
func _solveNQueensIsVal(qs []int, row int, column int) bool {
	for i := 0; i < row; i++ {
		if qs[i] == column {
			return false
		}
		if qs[i] == column-(row-i) {
			return false
		}
		if qs[i] == column+(row-i) {
			return false
		}
	}
	return true
}

// 生成结果
func _solveNQueensGeneralRet(qs []int) {
	n := len(qs)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		rowstr := make([]byte, n)
		for j := 0; j < n; j++ {
			if qs[i] == j {
				rowstr[j] = 'Q'
			} else {
				rowstr[j] = '.'
			}
		}
		res[i] = string(rowstr)
	}
	_solveNQueensRet = append(_solveNQueensRet, res)
	return
}
