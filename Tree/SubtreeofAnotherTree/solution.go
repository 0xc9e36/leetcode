/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 16:48
 */

package SubtreeofAnotherTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {


	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	return isSameStruct(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameStruct(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	return (s.Val == t.Val) && isSameStruct(s.Left, t.Left) && isSameStruct(s.Right, t.Right)
}