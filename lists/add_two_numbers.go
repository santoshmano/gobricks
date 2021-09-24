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
	// add the numbers and the remainder and keep creating new nodes for the new list
	// add a last node if remainder exists

	head := ListNode{0, nil}
	cur := &head
	var remainder int
	var sum int

	for (l1 != nil) || (l2 != nil) {
		x, y := 0, 0

		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		sum = x + y + remainder
		remainder = sum / 10
		sum = sum % 10

		node := ListNode{sum, nil}
		cur.Next = &node
		cur = &node

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if remainder != 0 {
		node := ListNode{remainder, nil}
		cur.Next = &node
		cur = &node
	}

	//fmt.Println(head.Next, head.Val, cur.Val, cur.Next)
	return head.Next
}
