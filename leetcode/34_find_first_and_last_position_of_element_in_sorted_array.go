package leetcode

// first first index where target is found, else -1
func binarySearchFirst(nums []int, target int) int {

	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	if start < len(nums) && nums[start] == target {
		return start
	}

	return -1
}

// first first index where target is found, else -1
func binarySearchLast(nums []int, target int) int {

	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if (end > -1) && (nums[end] == target) {
		return end
	}

	return -1
}

func searchRange(nums []int, target int) []int {

	firstPos := -1
	lastPos := -1

	firstPos = binarySearchFirst(nums, target)
	lastPos = binarySearchLast(nums, target)

	return []int{firstPos, lastPos}
}
