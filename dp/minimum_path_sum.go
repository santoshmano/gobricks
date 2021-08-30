package dp

import "fmt"

// https://leetcode.com/problems/minimum-path-sum/

// Recurrence Relations
//  minPS(m,n) = min(minPS(m-1, n), minPS(m, n-1)) + cost[m][n]
//
// T(n) = O(2^m+n)
// S(n) = O(m+n)
//
// m(1,1) = min(m(0,1), m(1,0)) + grid[1][1]

func minPathSum(grid [][]int) int {

	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	minPS := make([][]int, rows)
	for i := 0; i < rows; i++ {
		minPS[i] = make([]int, cols)
	}

	minPS[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		minPS[i][0] = minPS[i-1][0] + grid[i][0]
	}

	for i := 1; i < cols; i++ {
		minPS[0][i] = minPS[0][i-1] + grid[0][i]
	}

	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			minPS[i][j] = min(minPS[i-1][j], minPS[i][j-1]) + grid[i][j]
		}
	}
	fmt.Println(grid, minPS)

	return minPS[rows-1][cols-1]
}
