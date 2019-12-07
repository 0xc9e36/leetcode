/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 15:45
 */

package BalancedBinaryTree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := height(root.Left)
	r := height(root.Right)

	if l - r == 0 || l - r == 1 || r - l == 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := height(root.Left)

	r := height(root.Right)

	if l > r {
		return l + 1
	}

	return r + 1
}