/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-25 21:17
 */

package MergekSortedLists

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

//优先队列
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq) - 1
	old := *pq
	*pq = old[:n]
	return old[n]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func mergeKLists(lists []*ListNode) *ListNode {

	pq := &PriorityQueue{}
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(pq, lists[i])
	}

	heap.Init(pq)

	head := &ListNode{}
	prev := head
	//队列不为空
	for pq.Len() != 0 {
		x := heap.Pop(pq)
		node := x.(*ListNode)
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		node.Next = nil
		prev.Next = node
		prev = node
	}

	return head.Next
}
