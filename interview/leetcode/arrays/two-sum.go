package arrays

// i/p : []int, array of random integers
//       int,  target number
// o/p : []int, array of 2 indices whose values add up to target.

// scan array and find two numbers that add up to target
//
// observations
//      - if there are multiple such sets, return first set
//

// approach 1
//      for every number, scan the array and find if it has another
//      corresponding number that adds upto to target
// T(n) = O(n^2)
// S(n) = O(1)
//
func twoSumSlow(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}

		}
	}

	return []int{-1, -1}
}

// ex
//.  0. 1. 2.  3
// [ 2, 7, 11, 15] , tar = 9

// approach 2
//      add all numbers to a map (key=num, val=index)
//      scan the array, and for every num
//          check if the map has (target-num)
// T(n) = 2 scans
//      = O(n)
// S(n) = O(n)

// Note -
//      need to take care of duplicates
//      cannot use the same idx for both numbers

func twoSum2(nums []int, target int) []int {

	numMap := make(map[int][]int) // key=num, val=array of indices

	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[target-nums[i]]; !ok {
			numMap[nums[i]] = []int{i}
		} else {
			numMap[nums[i]] = append(numMap[nums[i]], i)
		}
	}

	for i := 0; i < len(nums); i++ {
		if arr, ok := numMap[target-nums[i]]; ok {

			for _, val := range arr {
				if val != i {
					return []int{i, val}
				}
			}
		}
	}

	// if not found
	return []int{-1, -1}
}

// approach 3
//     store diff in map instead of number, makes it easier.
//      addresses duplicates by default

func twoSum(nums []int, target int) []int {

	diffMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		if _, present := diffMap[nums[i]]; present {
			return []int{i, diffMap[nums[i]]}
		} else {
			diffMap[target-nums[i]] = i
		}
	}

	return []int{-1, -1}
}
