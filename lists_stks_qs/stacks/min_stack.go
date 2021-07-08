package stacks

// https://leetcode.com/problems/min-stack

/* quick leetcode solution, not performant or thread safe */
type MinStack struct {
	dataStk []int
	minStk  []int
	top     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}

func (this *MinStack) Push(x int) {

	this.dataStk = append(this.dataStk, x)

	if this.top <= -1 {
		this.minStk = append(this.minStk, x)
	} else if (this.top > -1) && (x < this.minStk[this.top]) {
		this.minStk = append(this.minStk, x)
	} else {
		this.minStk = append(this.minStk, this.minStk[this.top])
	}

	this.top += 1

}

func (this *MinStack) Pop() {
	if this.top < 0 {
		return
	}
	this.dataStk = this.dataStk[:this.top]
	this.minStk = this.minStk[:this.top]
	this.top -= 1
}

func (this *MinStack) Top() int {
	if this.top < 0 {
		return -1
	}

	return this.dataStk[this.top]
}

func (this *MinStack) GetMin() int {
	if this.top < 0 {
		return -1
	}

	return this.minStk[this.top]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
