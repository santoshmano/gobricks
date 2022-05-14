package design

// https://leetcode.com/problems/shuffle-an-array/
import "math/rand"
import "time"

type Solution struct {
	orig   []int
	result []int
}

func Constructor(nums []int) Solution {
	var sol Solution
	rand.Seed(time.Now().Unix())
	sol.orig = nums
	return sol
}

func (this *Solution) Reset() []int {
	return this.orig
}

func (this *Solution) Shuffle() []int {

	result := make([]int, len(this.orig))
	copy(result, this.orig)

	for i := len(result); i > 0; i-- {
		random := rand.Intn(i)
		result[random], result[i-1] = result[i-1], result[random]
	}
	return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
