package trees

//https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

func sortedArrayToBST(nums []int) *TreeNode {

	var _sorted func(start, end int) *TreeNode

	_sorted = func(start, end int) *TreeNode {

		if start > end {
			return nil
		}

		mid := start + (end-start)/2
		root := &TreeNode{
			Val:   nums[mid],
			Left:  _sorted(start, mid-1),
			Right: _sorted(mid+1, end),
		}
		return root
	}

	return _sorted(0, len(nums)-1)
}
