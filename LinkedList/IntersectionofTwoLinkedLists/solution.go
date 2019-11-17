/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-17 20:19
 */

package IntersectionofTwoLinkedLists


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


//巧妙的解法:

// a + c
// b + c   c 是公共部分
// 有 a + c + b = b + c + a 成立
//所以第二次循环时, 指针相等一定是末尾空节点或者相交处
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	nodeA, nodeB := headA, headB
	for nodeA != nodeB {
		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}
		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}
	return nodeB
}