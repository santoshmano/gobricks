package lists

// https://leetcode.com/problems/linked-list-cycle-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	// check if it has a cycle
	fast := head
	slow := head
	hasCycle := false

	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

	// get length of cycle
	temp := slow.Next
	cycleLen := 1

	for temp != slow {
		temp = temp.Next
		cycleLen += 1
	}

	// x + cycLen = listLen
	// In a list with cycle, the node after the end of the list is the beginning of the cycle
	// -- start a first pointer from head and run it for cycLen
	// -- after that start a second pointer from the head, along with continuing to
	//	  run the first pointer, and the point where both the pointer meet is the
	//	  beginning of the cycle.

	second := head
	first := head
	for ; cycleLen > 0; cycleLen-- {
		first = first.Next
	}

	for first != second {
		first = first.Next
		second = second.Next
	}

	return second
}
