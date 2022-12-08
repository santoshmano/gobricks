package trees

// https://leetcode.com/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var q []*TreeNode

	if root != nil {
		q = append(q, root)
	}

	for len(q) > 0 {
		var level []int
		count := len(q)

		for count > 0 {
			node := q[0]
			q = q[1:]
			count--

			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
