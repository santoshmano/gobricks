package searching

//https://leetcode.com/problems/fixed-point
func fixedPoint(arr []int) int {

	start := 0
	end := len(arr) - 1
	index := -1

	for start <= end {

		mid := start + (end-start)/2

		if arr[mid] == mid {
			index = mid
			end = mid - 1
		} else if arr[mid] < mid {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return index

}
