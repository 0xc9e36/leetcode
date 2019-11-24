/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 23:29
 */

package AddTwoNumbersII

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 3 4 2 7
	// 4 6 5
	// 7 0 8 7
	arr1, arr2 := []int{}, []int{}
	for l1 != nil || l2 != nil {
		if l1 != nil {
			arr1 = append([]int{l1.Val}, arr1...)
			l1 = l1.Next
		}
		if l2 != nil {
			arr2 = append([]int{l2.Val}, arr2...)
			l2 = l2.Next
		}
	}

	var head *ListNode
	up := 0
	i := 0
	for i< len(arr1) || i < len(arr2) {

		if i < len(arr1) {
			up += arr1[i]
		}

		if i < len(arr2) {
			up += arr2[i]
		}

		fmt.Println(up)

		node := &ListNode{
			Val:up % 10,
			Next: head,
		}
		i++
		up /= 10
		head = node
	}

	if up > 0 {
		node := &ListNode{
			Val:up % 10,
			Next: head,
		}
		head = node
	}

	return head
}



