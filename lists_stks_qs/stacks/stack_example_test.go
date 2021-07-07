package stacks_test

import (
	"fmt"
	"github.com/santoshmano/gobricks/lists/stacks"
)

func ExampleStack() {

	stk := stacks.NewStack(10)

	elems := []int{10, 20, 30, 40, 50}
	for _, val := range elems {
		fmt.Println("inserting val", val)
		stk.Push(val)

		if stk.Peek() != val {
			fmt.Println("head of stk is ", stk.Peek(), " it is supposed to be ", val)
		}
	}

	for i := 0; i < len(elems)/2; i++ {
		elems[i], elems[len(elems)-1-i] = elems[len(elems)-1-i], elems[i]
	}

	for _, val := range elems {
		if stk.Peek() != val {
			fmt.Println("head of stk is ", stk.Peek(), " it is supposed to be ", val)
		}

		if stk.Pop() != val {
			fmt.Println("pop failed")
		} else {
			fmt.Println("popped item", val)
		}

	}
	//Output:
	//inserting val 10
	//inserting val 20
	//inserting val 30
	//inserting val 40
	//inserting val 50
	//popped item 50
	//popped item 40
	//popped item 30
	//popped item 20
	//popped item 10
}
