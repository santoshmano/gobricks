package sorting

//https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {

	writeIndex := len(nums1) - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[writeIndex] = nums1[m-1]
			m--
		} else {
			nums1[writeIndex] = nums2[n-1]
			n--
		}
		writeIndex--
	}

	for m > 0 {
		nums1[writeIndex] = nums1[m-1]
		writeIndex--
		m--
	}

	for n > 0 {
		nums1[writeIndex] = nums2[n-1]
		writeIndex--
		n--
	}
}
