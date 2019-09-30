package RangeSumofBST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	v := root.Val
	if v < L {
		return rangeSumBST(root.Right, L, R)
	} else if v > R {
		return rangeSumBST(root.Left, L, R)
	}
	return v + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}
