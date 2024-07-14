package nqueens

func solveNQueens(n int) [][]string {
	var res [][]string

	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	var cols, diag1, diag2 = make([]bool, n), make([]bool, 2*n), make([]bool, 2*n)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			var solution []string
			for _, r := range board {
				solution = append(solution, string(r))
			}
			res = append(res, solution)
			return
		}

		for col := 0; col < n; col++ {
			if cols[col] || diag1[row+col] || diag2[row-col+n] {
				continue
			}
			board[row][col] = 'Q'
			cols[col], diag1[row+col], diag2[row-col+n] = true, true, true
			backtrack(row + 1)
			board[row][col] = '.'
			cols[col], diag1[row+col], diag2[row-col+n] = false, false, false
		}
	}

	backtrack(0)

	return res
}
