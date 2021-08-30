package dp

// https://leetcode.com/problems/unique-paths/

// i/p: int row, int col ; 1-indexed, m&n 1-100
// o/p: int unique paths
func uniquePathsRec(m int, n int) int {

	if m == 1 && n == 1 {
		return 1
	}

	if m <= 1 {
		return uniquePaths(m, n-1)
	} else if n <= 1 {
		return uniquePaths(m-1, n)
	} else {
		return uniquePaths(m-1, n) + uniquePaths(m, n-1)
	}

}

func uniquePaths(m int, n int) int {

	table := make([][]int, m)
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
	}

	for row := 0; row < m; row++ {
		table[row][0] = 1
	}

	for col := 0; col < n; col++ {
		table[0][col] = 1
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			table[row][col] = table[row-1][col] + table[row][col-1]
		}
	}

	return table[m-1][n-1]

}
