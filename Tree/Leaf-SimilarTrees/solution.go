/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-07 08:13
 */

package Leaf_SimilarTrees


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	nums1 := inOrder(root1)
	nums2 := inOrder(root2)

	if len(nums1) != len(nums2) {
		return false
	}

	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) != 0 {
			node := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if node.Left == nil && node.Right == nil {
				res = append(res, node.Val)
			}
			root = node.Right
		}
	}
	return res
}