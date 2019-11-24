/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-24 21:21
 */

package MinStack

type MinStack struct {
	stack1, stack2 []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {

	this.stack1 = append(this.stack1, x)

	if len(this.stack2) == 0 || x <= this.stack2[len(this.stack2) - 1] {
		this.stack2 = append(this.stack2, x)
	}
}


func (this *MinStack) Pop()  {
	v := this.stack1[len(this.stack1) - 1]
	this.stack1 =  this.stack1[:len(this.stack1) - 1]
	if v == this.GetMin() {
		this.stack2 =  this.stack2[:len(this.stack2) - 1]
	}
}


func (this *MinStack) Top() int {
	return this.stack1[len(this.stack1) - 1]
}


func (this *MinStack) GetMin() int {
	return this.stack2[len(this.stack2) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */