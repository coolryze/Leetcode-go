package sudokusolver

func solveSudoku(board [][]byte) {
	isValid := func(board [][]byte, row, col int, c byte) bool {
		for i := 0; i < 9; i++ {
			if board[row][i] == c || board[i][col] == c || board[(row/3)*3+i/3][(col/3)*3+i%3] == c {
				return false
			}
		}
		return true
	}

	var solve func(board [][]byte) bool
	solve = func(board [][]byte) bool {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] == '.' {
					for c := byte('1'); c <= '9'; c++ {
						if isValid(board, row, col, c) {
							board[row][col] = c
							if solve(board) {
								return true
							}
							board[row][col] = '.'
						}
					}
					return false
				}
			}
		}
		return true
	}

	solve(board)
}
