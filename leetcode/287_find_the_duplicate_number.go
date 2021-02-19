package leetcode

func findDuplicate(nums []int) int {

	for i := 0; i < len(nums); i++ {

		for (i + 1) != nums[i] {

			val := nums[i]
			if val == nums[val-1] {
				return val
			}

			nums[i], nums[val-1] = nums[val-1], nums[i]
		}
	}
	return -1
}
