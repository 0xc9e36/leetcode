/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 21:26
 */

package NextGreaterNodeInLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	list := []int{}
	for node := head; node != nil; node = node.Next {
		list = append(list, node.Val)
	}
	l, ans := len(list), make([]int, len(list))
	stack, top := make([]int, l), -1
	for i := 0; i < l; i++ {
		for top >= 0 && list[stack[top]] < list[i] {
			ans[stack[top]] = list[i]
			top--
		}
		top++
		stack[top] = i
	}
	return ans
}
