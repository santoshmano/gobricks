package sorting

/*
Divide, sort, merge two sorted arrays

  - Given a range of elems divide it into 2 equal(+/-1) halves

  - Repeat the above till there are two sorted halves
   (1/0 elem array being the smallest sorted array)

  - Having two independent sorted ranges, merge them together
    to form one big sorted range

	T(n) - O(nlogn)
	Average - O(nlogn)
	Worst - O(nlogn)
	Best - O(nlogn)
	S(n) - O(n)
	Stable Sort - Yes
*/
func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start int, end int) {

	if start >= end {
		return
	}

	mid := (start + end) / 2 // replace this with non-overflow expression
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)

	merge(nums, start, mid, end)
}

func merge(nums []int, start int, mid int, end int) {
	aux := make([]int, len(nums))
	copy(aux, nums)

	i := start
	left := start
	right := mid + 1
	for (left <= mid) && (right <= end) {

		if aux[left] <= aux[right] {
			nums[i] = aux[left]
			left += 1
		} else {
			nums[i] = aux[right]
			right += 1
		}
		i += 1
	}

	for left <= mid {
		nums[i] = aux[left]
		left += 1
		i += 1
	}

	for right <= end {
		nums[i] = aux[right]
		right += 1
		i += 1
	}
}
