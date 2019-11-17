/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 16:35
 */

package DeleteNodeinaLinkedList

func deleteNode(node *ListNode) {

	node.Val = node.Next.Val
	if node.Next.Next == nil {
		node.Next = nil
		return
	}

	node.Next = node.Next.Next

	return
}