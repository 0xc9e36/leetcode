/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-26 22:00
 */

package RotateList

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	l := 0
	current := head
	tail := head
	for current != nil {
		l++
		tail = current
		current = current.Next
	}

	if l == 0 {
		return head
	}

	k = k % l

	current = head
	for i := l - k; i >= 2; i-- {
		if i < 2 {
			break
		}
		current = current.Next
	}

	tail.Next = head
	head = current.Next
	current.Next = nil

	return head
}
