package arrays

func isValidSudoku(board [][]byte) bool {

	var bit int

	// TODO - check if we need to boundary checks

	// check rows
	for row := 0; row < len(board); row++ {
		bit = 0
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] != '.' {
				val := board[row][col] - '0'

				if (bit & (1 << val)) != 0 {
					return false
				}
				bit = bit | (1 << val)
			}
		}
	}
	// check cols

	for row := 0; row < len(board); row++ {
		bit = 0
		for col := 0; col < len(board[0]); col++ {
			if board[col][row] != '.' {
				val := board[col][row] - '0'

				if (bit & (1 << val)) != 0 {
					return false
				}
				bit = bit | (1 << val)
			}
		}
	}

	// check squares

	for rowS := 0; rowS < len(board); rowS += 3 {
		for colS := 0; colS < len(board[0]); colS += 3 {

			bit = 0
			for row := rowS; row < rowS+3; row++ {
				for col := colS; col < colS+3; col++ {
					if board[row][col] != '.' {
						val := board[row][col] - '0'

						if (bit & (1 << val)) != 0 {
							return false
						}
						bit = bit | (1 << val)
					}
				}
			}
		}
	}

	return true
}
