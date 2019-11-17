/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 14:14
 */

package MiddleoftheLinkedList


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//快慢指针
//奇数 1 2 3  return slow
//偶数 1 2 3 4  return slow.Next
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return slow
		}
		if fast.Next.Next == nil {
			return slow.Next
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}