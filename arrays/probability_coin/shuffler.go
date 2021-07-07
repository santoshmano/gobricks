package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Given an array with 52 unique integers in the range [0,51]
// 0 -> ace spades, 1 -> ace clubs, 2 -> ace diamonds, 3 -> ace hearts

// shuffler takes in an array a and shuffles the ordering of the elements
func shuffle(a []int) {
	rand.Seed(time.Now().Unix())
	for i := len(a) - 1; i >= 0; i-- {
		x := rand.Intn(i + 1)
		a[i], a[x] = a[x], a[i]
	}
}

func f20() int {

	return 0
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(a)
	shuffle(a)
	fmt.Println(a)
}
