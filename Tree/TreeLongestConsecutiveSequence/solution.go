/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-08 20:47
 */

package TreeLongestConsecutiveSequence

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	dfs(root, root.Val, 0, &ans)
	return ans
}

func dfs(root *TreeNode, parent, len int, ans *int) {

	if root == nil {
		return
	}
	if parent + 1 == root.Val {
		len++
	} else {
		len = 1
	}

	if root.Left != nil {
		dfs(root.Left, root.Val, len, ans)
	}

	if root.Right != nil {
		dfs(root.Right, root.Val, len, ans)
	}

	*ans = max(*ans, len)
	return
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}