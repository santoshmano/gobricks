package sorting

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

// counting sort technique, O(1) space, O(n) time
func findDisappearedNumbers(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	var elemsNotPresent []int

	for i, val := range nums {
		if val != i+1 {
			elemsNotPresent = append(elemsNotPresent, i+1)
		}
	}

	return elemsNotPresent
}

// O(n) time, O(n) space, quick soln
func findDisappearedNumbersSimple(nums []int) []int {

	m := make(map[int]int)

	// initialize the map
	for i := 1; i <= len(nums); i++ {
		m[i-1] = 0
	}

	for _, val := range nums {
		m[val-1] += 1
	}

	var elemsNotPresent []int

	for key, val := range m {
		if val == 0 {
			elemsNotPresent = append(elemsNotPresent, key+1)
		}
	}

	return elemsNotPresent
}
