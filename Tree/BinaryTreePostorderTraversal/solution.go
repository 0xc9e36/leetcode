/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-06 22:53
 */

package BinaryTreePostorderTraversal


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//二叉树迭代遍历
func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append([]int{node.Val}, res...)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return res
}

