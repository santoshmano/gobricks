package arrays

// approach 1
//  sort both arrays, traverse both arrays at
//  same time, skip over duplicates
//  nlogn + n

// approach 2
//  create a hashmap of nums1 with a state , not present, present 1, present 0
//  traverse nums2 and if present 1, add to intersection list and set present to 0

func intersect(nums1 []int, nums2 []int) []int {

	nums1Map := make(map[int]int)
	var result []int

	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] += 1
	}

	for i := 0; i < len(nums2); i++ {
		val, ok := nums1Map[nums2[i]]
		if ok && (val > 0) {
			result = append(result, nums2[i])
			nums1Map[nums2[i]] -= 1
		}
	}

	return result
}
