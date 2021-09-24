package lists

// https://leetcode.com/problems/add-two-numbers/submissions/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// traverse both the lists
	// add the numbers and the remainder and create new list
	// traverse the remainder of the longer list

	head := ListNode{0, nil}
	cur := &head
	var remainder int
	var sum int

	for l1 != nil && l2 != nil {

		sum = l1.Val + l2.Val + remainder
		remainder = sum / 10
		sum = sum % 10

		node := ListNode{sum, nil}
		cur.Next = &node
		cur = &node
		//fmt.Println(cur.Val, cur.Next, l1.Val, l2.Val, sum, remainder)
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum = l1.Val + remainder
		remainder = sum / 10
		sum = sum % 10

		node := ListNode{sum, nil}
		cur.Next = &node
		cur = &node
		//fmt.Println(cur.Val, cur.Next, l1.Val, sum, remainder)
		l1 = l1.Next
	}

	for l2 != nil {
		sum = l2.Val + remainder
		remainder = sum / 10
		sum = sum % 10

		node := ListNode{sum, nil}
		cur.Next = &node
		cur = &node
		//fmt.Println(cur.Val, cur.Next, l2.Val, sum, remainder)
		l2 = l2.Next
	}

	if remainder != 0 {
		node := ListNode{remainder, nil}
		cur.Next = &node
		cur = &node
	}

	//fmt.Println(head.Next, head.Val, cur.Val, cur.Next)
	return head.Next
}
