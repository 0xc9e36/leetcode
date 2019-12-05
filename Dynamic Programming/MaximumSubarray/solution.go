/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-05 11:34
 */

package MaximumSubarray

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	dp := make([]int, l, l)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < l; i++ {
		dp[i] = max(nums[i], dp[i - 1] + nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}