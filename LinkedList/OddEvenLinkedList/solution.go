/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-23 10:16
 */

package OddEvenLinkedList


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next
	oddPrev, evenPrev := odd, even
	head = head.Next.Next
	oddPrev.Next, evenPrev.Next = nil, nil
	for head != nil {
		oddPrev.Next = head
		oddPrev = head
		if head.Next != nil {
			evenPrev.Next = head.Next
			evenPrev = head.Next
			head = head.Next.Next
			evenPrev.Next = nil
			continue
		}
		head = head.Next
		oddPrev.Next = nil
	}
	oddPrev.Next = even
	return odd
}
