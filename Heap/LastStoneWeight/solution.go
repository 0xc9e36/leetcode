/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-30 14:01
 */

package LastStoneWeight

import (
	"container/heap"
	"fmt"
)

type Stones []int

func (s Stones) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s Stones) Len() int {
	return len(s)
}

func (s Stones) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *Stones) Pop() interface{} {
	n := len(*s) - 1
	old := *s
	*s = old[:n]
	return old[n]
}

func (s *Stones) Push(x interface{}) {
	*s = append(*s, x.(int))
}


func lastStoneWeight(stones []int) int {

	pq := Stones(stones)

	heap.Init(&pq)
	fmt.Println(pq)
	for pq.Len() >= 2 {
		v1 := heap.Pop(&pq).(int)
		v2 := heap.Pop(&pq).(int)
		if v1 > v2 {
			heap.Push(&pq, v1 - v2)
		} else {
			heap.Push(&pq, v2 - v1)
		}
	}

	x := heap.Pop(&pq)
	return x.(int)
}