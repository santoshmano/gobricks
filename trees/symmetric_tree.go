package trees

// https://leetcode.com/problems/symmetric-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return traverse(root.Left, root.Right)
}

func traverse(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 != nil && r2 != nil {
		if r1.Val != r2.Val {
			return false
		}
		return traverse(r1.Left, r2.Right) && traverse(r1.Right, r2.Left)
	}

	return false
}
