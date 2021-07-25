package trees

//https://leetcode.com/problems/maximum-depth-of-binary-tree/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDepth := 1 + maxDepth(root.Left)
	rightDepth := 1 + maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}
