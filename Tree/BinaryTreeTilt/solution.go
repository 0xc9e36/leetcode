/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 18:40
 */

package BinaryTreeTilt

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


//暴力解法
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	l := sum(root.Left)
	r := sum(root.Right)


	if l - r > 0 {
		return l - r + findTilt(root.Left) + findTilt(root.Right)
	} else {
		return r - l + findTilt(root.Left) + findTilt(root.Right)
	}
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := sum(root.Left)
	right := sum(root.Right)
	return root.Val + left + right
}