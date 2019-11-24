/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 22:41
 */

package LinkedListComponents


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
	lg := len(G)
	m := make(map[int]struct{}, lg)
	for i := 0; i < lg; i++ {
		m[G[i]] = struct{}{}
	}

	ans := 0
	for head != nil {
		node := head
		for node != nil {
			if _, ok := m[node.Val]; !ok {
				break
			}
			node = node.Next
		}
		if head != node {
			ans++
			head = node
		} else {
			head = head.Next
		}
	}
	return ans
}