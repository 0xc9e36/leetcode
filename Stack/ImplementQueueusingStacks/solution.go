/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-24 10:39
 */

package ImplementQueueusingStacks

type MyQueue struct {
	stack1 []int
	stack2 []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack1 : make([]int, 0),
		stack2 : make([]int, 0),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	for len(this.stack1) != 0 {
		this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
		this.stack1 = this.stack1[:len(this.stack1) - 1]
	}
	this.stack2 = append(this.stack2, x)

	for len(this.stack2) != 0 {
		this.stack1 = append(this.stack1, this.stack2[len(this.stack2) - 1])
		this.stack2 = this.stack2[:len(this.stack2) - 1]
	}
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	x := this.stack1[len(this.stack1) - 1]
	this.stack1 = this.stack1[:len(this.stack1) - 1]
	return x
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.stack1[len(this.stack1) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */