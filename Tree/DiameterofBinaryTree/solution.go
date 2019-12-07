/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 20:20
 */

package DiameterofBinaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {

	res := 0
	postOrder(root, &res)
	return res
}

func postOrder(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	l := postOrder(root.Left, res)
	r := postOrder(root.Right, res)

	if l + r > *res {
		*res = l + r
	}

	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}