package sorting

/*
Select the smallest element, place it at leftmost index

  - given a set of elements ranging (i:n) find the smallest element
   and swap it with element at i.

  - repeat the above for range (i+1:n)

  T(n) - O(n**2)
  S(n) - O(1)
  Stable - No
*/
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		// find the index of smallest element
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// place smallest element at index i
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func SelectionSortRec(nums []int) {
	selectionSortRec(nums, 0, len(nums)-1)
}

func selectionSortRec(nums []int, start int, end int) {

	if start >= end {
		return
	}

	minIndex := start
	for i := start + 1; i <= end; i++ {
		if nums[i] < nums[minIndex] {
			minIndex = i
		}
	}
	nums[start], nums[minIndex] = nums[minIndex], nums[start]

	selectionSortRec(nums, start+1, end)
}
