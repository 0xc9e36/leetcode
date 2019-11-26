/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-24 22:09
 */

package SlidingWindowMaximum

//单调队列
func maxSlidingWindow(nums []int, k int) []int {
	n, res := len(nums), []int{}
	if n == 0 {
		return res
	}
	queue := []int{}
	for i := 0; i < n; i++ {
		for len(queue) != 0 && i - queue[0] >= k {
			queue = queue[1:]
		}
		for len(queue) != 0 && nums[i] > nums[queue[len(queue) - 1]] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, i)
		res = append(res, nums[queue[0]])
	}
	return res[k - 1:]
}
