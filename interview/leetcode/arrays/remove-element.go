package arrays

// https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}

	return right + 1
}
