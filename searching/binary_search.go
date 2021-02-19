// https://leetcode.com/problems/binary-search/
package searching

// iterative binary search
func search(nums []int, target int) int {

	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

// recursive binary search
func search1(nums []int, target int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2

	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return search1(nums, target, mid+1, end)
	} else {
		return search1(nums, target, start, mid-1)
	}
}
