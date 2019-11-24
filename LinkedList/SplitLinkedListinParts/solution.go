/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-23 11:33
 */

package SplitLinkedListinParts

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	root := &ListNode{
		Next: head,
	}
	prev := root
	for head.Next != nil {
		swap := head.Next
		head.Next = swap.Next
		swap.Next = head
		prev.Next = swap
		head = head.Next
		prev = swap.Next
		if head == nil {
			break
		}
	}
	return root.Next
}