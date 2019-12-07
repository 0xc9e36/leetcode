/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 08:01
 */

package IncreasingOrderSearchTree



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	res := &TreeNode{}
	prev := res
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) != 0 {
			node := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			prev.Right = node
			prev = node
			prev.Left = nil
			root = node.Right
		}
	}
	return res.Right
}