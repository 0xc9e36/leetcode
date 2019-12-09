/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-08 19:09
 */

package InsertintoaBinarySearchTree


func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val: val,
		}
		return root
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}