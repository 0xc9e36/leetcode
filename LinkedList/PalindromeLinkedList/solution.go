/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 18:30
 */

package PalindromeLinkedList


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func isPalindrome(head *ListNode) bool {
	l := 0
	for node := head; node != nil; node = node.Next {
		l++
	}
	if l < 2 {
		return true
	}

	var left, right *ListNode
	left = head
	head = head.Next
	for i := 0; i < (l / 2) - 1; i++ {
		v := head.Next
		head.Next = left
		left = head
		head = v
	}

	if l % 2 == 0 {
		right = head
	} else {
		right = head.Next
	}

	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}