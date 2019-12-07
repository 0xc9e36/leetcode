/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 09:28
 */

package LongestUnivaluePath


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, ans)
	r := dfs(root.Right, ans)

	if root.Left != nil && root.Left.Val == root.Val {
		l++
	} else {
		l = 0
	}

	if root.Right != nil && root.Right.Val == root.Val {
		r++
	} else {
		r = 0
	}

	*ans = max(*ans, l + r)
	return max(l, r)
}


func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}