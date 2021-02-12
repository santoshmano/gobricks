package leetcode

func singleNonDuplicate(nums []int) int {

	start := 0
	end := len(nums) - 1

	// Useful to do boundary checks in such problems so that
	// you can avoid doing it over and over in the loop

	if len(nums) == 1 {
		return nums[start]
	} else if nums[start] != nums[start+1] {
		return nums[start]
	} else if nums[end] != nums[end-1] {
		return nums[end]
	}

	// iter len_window  start   end     mid
	// 1    odd         even    even    odd
	// 2    odd         even    even    even
	// 3    even        even    odd     even

	// inferences:

	// mid can be odd/even
	// if window is odd, it has the target (this is obvious)

	// the window, left of target has even-odd-indexed pairs of numbers
	//  - every pair is on
	// the window, right of target has odd-even-indexed pairs of numbers

	for start <= end {
		mid := start + (end-start)/2

		if mid%2 == 0 {
			// mid is even
			if nums[mid] == nums[mid+1] {
				// target is on right window
				start = mid + 1
			} else if mid > 0 && nums[mid] != nums[mid-1] {
				// return mid
				return nums[mid]
			} else {
				end = mid - 1
			}
		} else if mid%2 == 1 {
			// mid is odd
			if nums[mid] == nums[mid-1] {
				// target is on right
				start = mid + 1
			} else if mid < len(nums)-1 && nums[mid] != nums[mid+1] {
				return nums[mid]
			} else {
				end = mid - 1
			}
		}
	}

	return nums[start]
}
