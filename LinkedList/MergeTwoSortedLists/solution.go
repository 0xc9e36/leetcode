/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 16:46
 */

package MergeTwoSortedLists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := new(ListNode)


	if l2.Val < l1.Val {
		t := l2.Next
		l2.Next = l1
		l1 = l2
		l2 = t

	}
	head.Next = l1

	for l1 != nil && l2 != nil {

		if l1.Next == nil {
			l1.Next = l2
			break
		}

		if l1.Next.Val <= l2.Val {
			l1 = l1.Next
		} else {
			t := l1.Next
			v := l2.Next
			l1.Next = l2
			l1.Next.Next = t
			l2 = v
		}
	}

	return head.Next
}


//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//
//
//	if l1 == nil {
//		return l2
//	}
//
//	if l2 == nil {
//		return l1
//	}
//
//
//	//1 2
//	if l1.Val < l2.Val {
//		l1.Next = mergeTwoLists(l1.Next, l2)
//		return l1
//	} else {
//		l2.Next = mergeTwoLists(l1, l2.Next)
//		return l2
//	}
//
//}