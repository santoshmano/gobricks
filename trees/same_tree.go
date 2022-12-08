package trees

// https://leetcode.com/problems/same-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Every node, check if val is same, if they have
	//     equivalent Left and Right pointers

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if isSameTree(p.Left, q.Left) == false {
		return false
	}

	if isSameTree(p.Right, q.Right) == false {
		return false
	}

	return true

}
