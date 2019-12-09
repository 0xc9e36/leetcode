/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 22:26
 */

package SymmetricTree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil{
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}

	return isSymmetricHelp(root.Left, root.Right)
}


func isSymmetricHelp(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil {
		return false
	}

	return (r1.Val == r2.Val) && isSymmetricHelp(r1.Left, r2.Right) && isSymmetricHelp(r1.Right, r2.Left)
}

