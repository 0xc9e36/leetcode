/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-05 12:37
 */

package RangeSumQuery_Immutable

type NumArray struct {
	dp []int
	nums []int
}


func Constructor(nums []int) NumArray {
	l := len(nums)
	if l == 0 {
		return NumArray{}
	}
	dp := make([]int, l, l)
	dp[0] = nums[0]
	for i := 1; i < l; i++ {
		dp[i] = dp[i - 1] + nums[i]
	}

	return NumArray{
		dp: dp,
		nums: nums,
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.dp[j] - this.dp[i] + this.nums[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */