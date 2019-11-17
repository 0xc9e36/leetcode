/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 14:28
 */

package ReverseLinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1 2 3 4 5
// 1 2 3 4 5
// 链表插入, 头插法
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		v := curr.Next
		curr.Next = prev
		prev = curr
		curr = v
	}
	return prev
}