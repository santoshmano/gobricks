package arrays

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	// have a read & write index
	read := 1
	write := 1
	for read < len(nums) {
		if nums[read] != nums[read-1] {
			nums[write] = nums[read]
			write++
		}
		read++
	}

	return write
}
