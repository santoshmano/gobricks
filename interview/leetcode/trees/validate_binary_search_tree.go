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

import "fmt"

/*
2^3 = 2*2*2 = 8 == 1<<3
1<<3 = b(1000) = 8
*/

// BST is valid, if
// - nil or single node
// - max value in the left subtree < root.val
// - min value in right subtree > root.val
func isValidBST(root *TreeNode) bool {

	var _isBST func(root *TreeNode, max int, min int) bool

	_isBST = func(root *TreeNode, max int, min int) bool {
		if root == nil {
			return true
		}
		if root.Val < max && root.Val > min {
			// true conditions
			return _isBST(root.Left, root.Val, min) && _isBST(root.Right, max, root.Val)
		}
		return false
	}

	return _isBST(root, 1<<31, -1<<31-1)
}
