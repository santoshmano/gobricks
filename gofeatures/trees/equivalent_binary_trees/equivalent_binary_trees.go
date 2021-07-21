package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRecursive(t, ch)

	// close the channel
	close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	WalkRecursive(t.Left, ch)
	ch <- t.Value
	WalkRecursive(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1Ch := make(chan int)
	t2Ch := make(chan int)

	go Walk(t1, t1Ch)
	go Walk(t2, t2Ch)

	for {
		v1, v1ok := <-t1Ch
		v2, v2ok := <-t2Ch
		if v1ok && v2ok && v1 != v2 {
			return true
		}

		if !v1ok || !v2ok {
			return true
		}
	}
	return true
}

func main() {

	fmt.Println("=============")
	fmt.Println("Positive Case")
	fmt.Println("=============")
	for _, i := range []int{2, 200, 0} {
		fmt.Println(i)
		t1 := tree.New(i)
		t2 := tree.New(i)
		fmt.Println("Tree1:", t1.String())
		fmt.Println("Tree2:", t2.String())

		if Same(t1, t2) == true {
			fmt.Println("T1 and T2 are equal")
		} else {
			fmt.Println("Same check failed")
		}
	}

	fmt.Println("=============")
	fmt.Println("Negative Case")
	fmt.Println("=============")
	for _, i := range []int{10, 20} {
		t3 := tree.New(i)
		t4 := tree.New(i / 2)
		fmt.Println(i, i/2)
		fmt.Println("Tree1:", t3.String())
		fmt.Println("Tree2:", t4.String())

		if Same(t3, t4) == false {
			fmt.Println("T1 and T2 are not equal")
		} else {
			fmt.Println("Same check failed")
		}
	}
}
