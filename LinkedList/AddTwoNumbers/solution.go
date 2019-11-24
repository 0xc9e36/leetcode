/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 22:54
 */

package AddTwoNumbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	up := 0
	head := new(ListNode)
	head.Next = l1
	prev := head
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + up
		l1.Val = val % 10
		prev = l1
		up = val / 10
		l1, l2 = l1.Next, l2.Next
	}


	for l1 != nil {
		val := l1.Val + up
		l1.Val = val % 10
		up = val / 10
		prev = l1
		l1 = l1.Next
	}

	if l2 != nil {
		prev.Next = l2
		for l2 != nil {
			val := l2.Val + up
			l2.Val = val % 10
			up = val / 10
			prev = l2
			l2 = l2.Next
		}
	}


	if l1 == nil && l2 == nil && up != 0 {
		prev.Next = &ListNode{
			Val: up,
		}
	}
	return head.Next
}
