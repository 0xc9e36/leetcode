/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 16:28
 */

package SecondMinimumNodeInaBinaryTree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	first, second := math.MaxInt64, math.MaxInt64
	dfs(root, &first, &second)
	if second == math.MaxInt64 {
		return -1
	}
	return second
}

func dfs(root *TreeNode, first, second *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if root.Val < *first {
			*second =  *first
			*first = root.Val
		} else if root.Val > *first && root.Val < *second {
			*second = root.Val
		}
	}

	dfs(root.Left, first, second)
	dfs(root.Right, first, second)
}
