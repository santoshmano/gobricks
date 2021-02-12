package sorting

import (
	"testing"
)

func arrayEqual(in []int, want []int) bool {
	if len(in) != len(want) {
		return false
	}

	for i, val := range in {
		if want[i] != val {
			return false
		}
	}

	return true
}

func genSortingTestData() []struct {
	orig  []int
	input []int
	want  []int
} {
	return []struct {
		orig  []int
		input []int
		want  []int
	}{
		{[]int{9, 10, -1, 0, 101, 33}, []int{9, 10, -1, 0, 101, 33}, []int{-1, 0, 9, 10, 33, 101}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 3, 4, 6, 7}, []int{1, 3, 4, 6, 7}, []int{1, 3, 4, 6, 7}},
		{[]int{1}, []int{1}, []int{1}},
		{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{-10, -12134, -4252524, -1, 0, -1}, []int{-10, -12134, -4252524, -1, 0, -1}, []int{-4252524, -12134, -10, -1, -1, 0}},
		{[]int{}, []int{}, []int{}},
		{nil, nil, nil}}
}

func TestBubbleSortInt(t *testing.T) {
	for _, testcase := range genSortingTestData() {
		BubbleSort(testcase.input)
		if arrayEqual(testcase.input, testcase.want) == false {
			t.Errorf(`BubbleSort(orig input:%v) generated: %v  expected: %v`,
				testcase.orig,
				testcase.input,
				testcase.want)
		} else {
			t.Log(`test success`)
		}
	}
}

func TestBubbleSortRecInt(t *testing.T) {
	for _, testcase := range genSortingTestData() {
		BubbleSortRec(testcase.input)
		if arrayEqual(testcase.input, testcase.want) == false {
			t.Errorf(`BubbleSort(orig input:%v) generated: %v  expected: %v`,
				testcase.orig,
				testcase.input,
				testcase.want)
		} else {
			t.Log(`test success`)
		}
	}
}

func TestSelectionSortInt(t *testing.T) {
	for _, testcase := range genSortingTestData() {
		SelectionSort(testcase.input)
		if arrayEqual(testcase.input, testcase.want) == false {
			t.Errorf(`BubbleSort(orig input:%v) generated: %v  expected: %v`,
				testcase.orig,
				testcase.input,
				testcase.want)
		} else {
			t.Log(`test success`)
		}
	}
}

func TestSelectionSortRecInt(t *testing.T) {
	for _, testcase := range genSortingTestData() {
		SelectionSortRec(testcase.input)
		if arrayEqual(testcase.input, testcase.want) == false {
			t.Errorf(`BubbleSort(orig input:%v) generated: %v  expected: %v`,
				testcase.orig,
				testcase.input,
				testcase.want)
		} else {
			t.Log(`test success`)
		}
	}
}
func TestInsertionSortInt(t *testing.T) {
	for _, testcase := range genSortingTestData() {
		InsertionSort(testcase.input)
		if arrayEqual(testcase.input, testcase.want) == false {
			t.Errorf(`BubbleSort(orig input:%v) generated: %v  expected: %v`,
				testcase.orig,
				testcase.input,
				testcase.want)
		} else {
			t.Log(`test success`)
		}
	}
}
func TestQuickSort(t *testing.T) {
	for _, testcase := range genSortingTestData() {
		QuickSort(testcase.input)
		if arrayEqual(testcase.input, testcase.want) == false {
			t.Errorf(`QuickSort(orig input:%v) generated: %v  expected: %v`,
				testcase.orig,
				testcase.input,
				testcase.want)
		} else {
			t.Log(`test success`)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, testcase := range genSortingTestData() {
		MergeSort(testcase.input)
		if arrayEqual(testcase.input, testcase.want) == false {
			t.Errorf(`MergeSort(orig input:%v) generated: %v expected: %v`,
				testcase.orig,
				testcase.input,
				testcase.want)
		} else {
			t.Log(`test success`)
		}
	}
}
