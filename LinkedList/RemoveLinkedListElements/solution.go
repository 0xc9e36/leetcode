/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 19:55
 */

package RemoveLinkedListElements


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	root := new(ListNode)
	root.Next = head
	prev := root
	for head != nil {
		if head.Val != val {
			prev.Next = head
			prev = head
		}
		head = head.Next
	}
	prev.Next = nil
	return root.Next
}