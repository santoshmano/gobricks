package trees

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// observations
//  - every complete level has 2^h nodes
//  - last level have 1 to 2^h nodes

// https://leetcode.com/problems/check-completeness-of-a-binary-tree/submissions/

// i/p - root of a binary tree
// o/p - bool, true if complete, false if not complete

// approach 1
//    traverse tree using level order,
//    count the nodes in every level and
//    confirm it has 2^h nodes, and if not confirm it is last level.
//      lvl1 = 1<<0 node  == 2^0
//      lvl2 = 1<<1 node  == 2^1
//      lvl3 = 1<<2 nodes == 2^2
//      lvlh = 1<<(h-1) nodes

func isCompleteTree(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var q []*TreeNode
	q = append(q, root)

	lvl := 0
	var isComplete bool
	var node *TreeNode
	var hole bool

	for len(q) > 0 {

		nodeCount := len(q)
		if nodeCount < (1 << lvl) {
			isComplete = false
		} else {
			isComplete = true
		}

		hole = false
		fmt.Println(nodeCount, isComplete)
		for nodeCount > 0 {
			node, q = q[0], q[1:] //pop
			nodeCount--

			if node.Left != nil {
				if hole {
					return false
				}
				q = append(q, node.Left)
			} else {
				hole = true
			}

			if node.Right != nil {
				if hole {
					return false
				}
				q = append(q, node.Right)
			} else {
				hole = true
			}

		}

		if !isComplete && len(q) != 0 {
			return false
		}

		lvl++

	}

	return true

}
