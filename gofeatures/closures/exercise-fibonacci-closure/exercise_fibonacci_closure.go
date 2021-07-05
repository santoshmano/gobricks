// https://tour.golang.org/moretypes/26

package main

import "fmt"

func fibonacci() func() int {
	first, second := 0, 1

	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci1() func() int {
	fibArray := []int{}

	return func() int {
		switch len(fibArray) {
		case 0:
			fibArray = append(fibArray, 0)
			return 0
		case 1:
			fibArray = append(fibArray, 1)
			return 1
		default:
			l1 := fibArray[len(fibArray)-1]
			l2 := fibArray[len(fibArray)-2]

			fibArray = append(fibArray, l1+l2)
			return l1 + l2
		}
		return -1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
