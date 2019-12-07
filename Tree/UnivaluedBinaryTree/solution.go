/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-06 23:40
 */

package UnivaluedBinaryTree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil {
		return (root.Val == root.Right.Val) && isUnivalTree(root.Right)
	}

	if root.Right == nil {
		return (root.Val == root.Left.Val) && isUnivalTree(root.Left)
	}

	return root.Val == root.Left.Val && root.Val == root.Right.Val && isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

