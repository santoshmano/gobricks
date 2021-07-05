package sorting

/*
Pick a pivot, move smaller to left and bigger to right

  - pick a random elem at pos i as a pivot

  - find the correct pos for it in the array by  placing all the
    smaller elems to its left and  bigger elems to the right.

  - repeat the above for range (0,i-1) & (i+1,n)

	T(n) - O(nlogn)
	S(n) - O(1)
	Stable Sort - No
*/
func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}

	pivot := partition(nums, start, end)
	quickSort(nums, start, pivot-1)
	quickSort(nums, pivot+1, end)
}

func partition(nums []int, start int, end int) int {
	pivot := start
	left := start + 1
	right := end

	for left <= right {
		if nums[left] <= nums[pivot] {
			left += 1
		} else if nums[right] >= nums[pivot] {
			right -= 1
		} else {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	nums[right], nums[pivot] = nums[pivot], nums[right]
	return right
}
