package arrays

// nxn grid, square
// rotate right by 90deg
/*

  0    1    2    3
 0,0  0,1  0,2  0,3  0
 1,0  1,1  1,2  1,3  1
 2,0  2,1  2,2  2,3  2
 3,0  3,1  3,2  3,3  3

 r,c			r, n-r-1



n-r-1, r		n-r-1, n-c-1

*/

// approach 1 - allocate a new 2d array and place it in it the position
// S(n) - O(n*n)

// approach 2 - allocate an array of size n and work 1 row/col at a time
// S(n) - O(n*n)

// approach 3 - rotate one m,n location at a time and place it in its position
//	- each iteration we are rotating 4 cells

func rotate(matrix [][]int) {
	n := len(matrix)
	for r := 0; r < (n/2 + n%2); r++ {
		for c := r; c < n/2; c++ {
			temp := matrix[r][c]

			matrix[r][c] = matrix[n-r-1][r]
			matrix[n-r-1][r] = matrix[n-r-1][n-c-1]
			matrix[n-r-1][n-c-1] = matrix[c][n-r-1]
			matrix[r][n-r-1] = temp
		}
	}
	return
}
