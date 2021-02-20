package main

import "fmt"

func genTestData() []struct {
	inputArr  []int
	inputKey  int
	outputIdx int
} {
	return []struct {
		inputArr  []int
		inputKey  int
		outputIdx int
	}{
		{[]int{}, 5, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, -1, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 10, -1},
	}
}

func binarySearch(arr []int, key int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2

	if key == arr[mid] {
		return mid
	} else if key < arr[mid] {
		return binarySearch(arr, key, left, mid-1)
	} else {
		return binarySearch(arr, key, mid+1, right)
	}
}

func BinarySearch(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}
	return binarySearch(arr, key, 0, len(arr)-1)
}

func main() {
	for _, in := range genTestData() {
		if BinarySearch(in.inputArr, in.inputKey) != in.outputIdx {
			fmt.Println("Binary search error")
		}
		fmt.Println("Binary search succeeded")
	}
}
