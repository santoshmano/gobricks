package leetcode

func searchMatrix(matrix [][]int, target int) bool {

	rows := len(matrix)
	cols := len(matrix[0])

	// identify the row where target could exist
	start := 0
	end := rows - 1
	targetRow := -1

	for start <= end {
		mid := start + (end-start)/2

		if target >= matrix[mid][0] {
			if target <= matrix[mid][cols-1] {
				targetRow = mid
				break
			} else {
				start = mid + 1
			}
		} else {
			end = mid - 1
		}
	}

	if targetRow == -1 {
		return false
	}

	// identify the col where target could exist

	start = 0
	end = cols - 1

	for start <= end {

		mid := start + (end-start)/2

		if target == matrix[targetRow][mid] {
			return true
		} else if matrix[targetRow][mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}
