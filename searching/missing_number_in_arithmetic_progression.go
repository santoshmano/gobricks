package searching

// https://leetcode.com/problems/missing-number-in-arithmetic-progression

func missingNumber(arr []int) int {

	if len(arr) <= 2 {
		return -1
	}

	// ap = n, n+1, n+2d, n+3d.....

	start := 1
	end := len(arr) - 1

	constant := (arr[len(arr)-1] - arr[0]) / len(arr)

	for start <= end {

		mid := start + (end-start)/2

		expected := arr[0] + mid*constant

		if arr[mid] == expected {
			start = mid + 1
		} else {
			if (arr[mid] - arr[mid-1]) != constant {
				return arr[mid] - constant
			}
			end = mid - 1
		}

	}

	return arr[end]

}
