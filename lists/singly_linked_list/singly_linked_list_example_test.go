package singly_linked_list_test

import (
	"fmt"
	lists "github.com/santoshmano/gobricks/lists_stks_qs/singly_linked_list"
)

func ExampleSinglyLinkedList() {
	l := lists.NewSinglyLinkedList()
	l.InsertVal(5)
	l.InsertVal(4)
	l.InsertVal(3)
	l.InsertVal(2)
	l.InsertVal(1)
	l.InsertVal("Numbers 1-5")
	for n := l.Head(); n != nil; n = n.Next {
		fmt.Println(n.Val)
	}

	l.DeleteVal(3)
	l.DeleteVal(1)
	l.DeleteVal(10)
	for n := l.Head(); n != nil; n = n.Next {
		fmt.Println(n.Val)
	}

	for _, v := range []interface{}{5, 3, 4, "Numbers 1-5"} {
		if l.FindVal(v) == v {
			fmt.Println(v, "Found")
		} else {
			fmt.Println(v, "NotFound")
		}
	}

	l.DeleteVal(2)
	l.DeleteVal(4)
	l.DeleteVal(5)
	for n := l.Head(); n != nil; n = n.Next {
		fmt.Println(n.Val)
	}

	l.DeleteVal("Numbers 1-5")
	for n := l.Head(); n != nil; n = n.Next {
		fmt.Println(n.Val)
	}

	fmt.Println(l.FindVal("Numbers 1-5"))

	//Output:
	//Numbers 1-5
	//1
	//2
	//3
	//4
	//5
	//Numbers 1-5
	//2
	//4
	//5
	//5 Found
	//3 NotFound
	//4 Found
	//Numbers 1-5 Found
	//Numbers 1-5
	//<nil>
}

func ExampleSinglyLinkedList_HasCycle() {
	l := lists.NewSinglyLinkedList()
	nodes := [5]*lists.SinglyLinkedNode{}

	fmt.Println("l has cycle?", l.HasCycle())

	for i, _ := range nodes {
		nodes[i] = &lists.SinglyLinkedNode{}
		nodes[i].Val = i
		l.InsertNode(nodes[i])
	}

	for n := l.Head(); n != nil; n = n.Next {
		fmt.Println(n.Val)
	}

	fmt.Println("l has cycle?", l.HasCycle())

	for n := 0; n < 5; n += 1 {
		//	fmt.Println(nodes[n], n)
	}

	nodes[0].Next = nodes[2]
	fmt.Println("l has cycle?", l.HasCycle())

	nodes[0].Next = nil
	fmt.Println("l has cycle?", l.HasCycle())

	//Output:
	//l has cycle? false
	//4
	//3
	//2
	//1
	//0
	//l has cycle? false
	//l has cycle? true
	//l has cycle? false
}
