package arrays

// https://leetcode.com/problems/contains-duplicate/solution/

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool)

	for _, num := range nums {
		_, ok := numsMap[num]
		if ok {
			return true
		}
		numsMap[num] = true
	}

	return false
}
