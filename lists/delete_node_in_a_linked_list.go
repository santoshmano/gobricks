package lists

// https://leetcode.com/problems/delete-node-in-a-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	// tricky hack
	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	node.Val = node.Next.Val
	node.Next = nil
}
