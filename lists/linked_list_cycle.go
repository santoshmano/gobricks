package lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next      /* slow runner 1 step at a time */
		fast = fast.Next.Next /* fast runner 2 steps at a time */

		/* if they meet eventually, there is a cycle */
		if slow == fast {
			return true
		}
	}

	return false
}

// Implementation with Maps, O(n) space complexity
func hasCycle1(head *ListNode) bool {

	nodeMap := make(map[*ListNode]int)

	for head != nil {
		_, present := nodeMap[head]
		if present {
			return true
		}
		nodeMap[head] = 1
		head = head.Next
	}

	return false
}
