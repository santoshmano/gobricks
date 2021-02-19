package sorting

// https://leetcode.com/problems/find-all-duplicates-in-an-array/

func findDuplicates(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				break
			} else {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			}
		}
	}

	var duplist []int

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			duplist = append(duplist, nums[i])
		}
	}

	return duplist
}
