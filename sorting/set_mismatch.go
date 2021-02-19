package sorting

// https://leetcode.com/problems/set-mismatch/

func findErrorNums(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				break
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return []int{nums[i], i + 1}
		}
	}
	return []int{-1, -1}
}
