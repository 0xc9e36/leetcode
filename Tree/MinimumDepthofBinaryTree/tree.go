/**
 * @Author: wei.tan
 * @Description:
 * @File:  tree
 * @Version: 1.0.0
 * @Date: 2019-12-07 11:26
 */

package MinimumDepthofBinaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		return max(l, r) + 1
	}

	return min(l, r) + 1
}


func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}