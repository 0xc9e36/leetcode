/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 22:14
 */

package BinaryTreeInorderTraversal


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) != 0 {
			node := stack[len(stack) - 1]
			res = append(res, node.Val)
			stack = stack[:len(stack) - 1]
			root = node.Right
		}
	}
	return res
}

