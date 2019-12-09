/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 20:23
 */

package BinaryTreePaths

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	l := binaryTreePaths(root.Left)
	r := binaryTreePaths(root.Right)

	for i := 0; i < len(l); i++ {
		l[i] = fmt.Sprintf("%d->%s", root.Val, l[i])
	}

	for i := 0; i < len(r); i++ {
		r[i] = fmt.Sprintf("%d->%s", root.Val, r[i])
	}

	res := append(l, r...)
	return res
}
