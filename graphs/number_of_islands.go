package main

func numIslands(grid [][]byte) int {

	// This problem can be solved by counting of number of
	// connected components approach
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	// create a 2d array to note the visited rows,cols
	// this is needed if we should not mutate the input grid
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	var dfs func(row, col int)

	dfs = func(row, col int) {

		visited[row][col] = true
		if grid[row][col] == '0' {
			return
		}

		if row-1 >= 0 && visited[row-1][col] == false {
			dfs(row-1, col)
		}
		if row+1 < rows && visited[row+1][col] == false {
			dfs(row+1, col)
		}
		if col-1 >= 0 && visited[row][col-1] == false {
			dfs(row, col-1)
		}
		if col+1 < cols && visited[row][col+1] == false {
			dfs(row, col+1)
		}
	}

	islands := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && visited[i][j] == false {
				islands += 1
				dfs(i, j)
			}
		}
	}

	return islands
}
