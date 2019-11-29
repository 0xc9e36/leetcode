/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-26 21:47
 */

package ReverseNodesink_Group

type ListNode struct {
	Val  int
	Next *ListNode
}

//虽然通过了, 但是这种写法太丑陋
func reverseKGroup(head *ListNode, k int) *ListNode {

	root := &ListNode{
		Next: head,
	}
	prev := root
	for head != nil {
		i := 0
		node := head
		p := head
		for i < k {
			p = head
			head = head.Next
			i++
			if head == nil {
				break
			}
		}

		if i != k {
			break
		}

		var next *ListNode
		if head != nil {
			next = head
		}
		p.Next = nil
		newHead, newTail := reverse(node)

		prev.Next = newHead
		newTail.Next = next

		prev = newTail
	}

	return root.Next
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	tail := head
	var prev *ListNode
	for head != nil {
		t := head.Next
		head.Next = prev
		prev = head
		head = t
	}
	return prev, tail
}
