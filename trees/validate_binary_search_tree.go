package trees

import "fmt"

//https://leetcode.com/problems/validate-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// valid BST if
	// -- nil or single node
	// -- max value in the left subtree < root.val
	// -- min value in right subtree > root.val

	return _isValidBST(root, -1<<63, 1<<63-1)
}

func _isValidBST(root *TreeNode, min int, max int) bool {

	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		fmt.Println(root.Val, min, max)
		return false
	}

	return _isValidBST(root.Left, min, root.Val) && _isValidBST(root.Right, root.Val, max)
}
