/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-08 09:48
 */

package PathSumII

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{{root.Val}}
		} else {
			return [][]int{}
		}
	}

	l := pathSum(root.Left, sum - root.Val)
	r := pathSum(root.Right, sum - root.Val)

	for i := 0; i < len(l); i++ {
		l[i] = append([]int{root.Val}, l[i]...)
	}

	for i := 0; i < len(r); i++ {
		r[i] = append([]int{root.Val}, r[i]...)
	}

	res := append(l, r...)
	return res
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func pathSum(root *TreeNode, sum int) [][]int {
//	path := []int{}
//	ans := [][]int{}
//	return pathSumHelp(root, sum, path, ans)
//}
//
//
//func pathSumHelp(root *TreeNode, sum int, path []int, ans [][]int) [][]int {
//	if root == nil {
//		return ans
//	}
//
//	sum -= root.Val
//	path = append(path, root.Val)
//
//	if root.Left == nil && root.Right == nil {
//		if sum == 0 {
//			p := make([]int, len(path))
//			copy(p, path)
//			ans = append(ans, p)
//		}
//		return ans
//	}
//
//	if root.Left != nil {
//		ans = pathSumHelp(root.Left, sum, path, ans)
//	}
//
//	if root.Right != nil {
//		ans = pathSumHelp(root.Right, sum, path, ans)
//	}
//
//	fmt.Println(ans)
//	return ans
//}
//
