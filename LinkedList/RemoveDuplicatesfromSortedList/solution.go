/**
 * @Author: wei.tan
 * @Description:
 * @File:  ssolution
 * @Version: 1.0.0
 * @Date: 2019-11-17 17:41
 */

package RemoveDuplicatesfromSortedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func deleteDuplicates(head *ListNode) *ListNode {
//	m := make(map[int]struct{}, 0)
//	root := new(ListNode)
//	root.Next = head
//	prev := root
//	for head != nil {
//		_, ok := m[head.Val]
//		if !ok {
//			prev.Next = head
//			prev = head
//			m[head.Val] = struct{}{}
//			head = head.Next
//			prev.Next = nil
//			continue
//		}
//
//		head = head.Next
//	}
//	return root.Next
//}



/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	root := &ListNode{
		Next: head,
	}
	prev := head
	for head != nil {
		if prev.Val != head.Val {
			prev.Next = head
			prev = head
		}
		head = head.Next
		prev.Next = nil
	}
	return root.Next
}