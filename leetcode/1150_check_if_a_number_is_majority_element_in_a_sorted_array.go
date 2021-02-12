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

func isMajorityElement(nums []int, target int) bool {

	firstIndex := binarySearchFirst(nums, target)

	if firstIndex == -1 {
		return false
	}

	lastIndex := firstIndex + len(nums)/2

	if lastIndex < len(nums) && nums[lastIndex] == target {
		return true
	}

	return false
}
