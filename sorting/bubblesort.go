package sorting

/*
Bubble the smallest element to leftmost unsorted position, repeat

  -	given a set of elements ranging (i, n)

  - repeat the above for range (i+1, n)

	T(n) - O(n**2)
	S(n) - O(1)
	Stable sort - Yes
*/
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		// bubble the smallest element to the ith index
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func BubbleSortRec(nums []int) {
	bubbleSortRec(nums, 0, len(nums)-1)
}

func bubbleSortRec(nums []int, start int, end int) {
	if start >= end {
		return
	}

	for i := end; i > start; i-- {
		if nums[i] < nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}

	bubbleSortRec(nums, start+1, end)
}
