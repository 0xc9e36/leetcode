/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-22 21:31
 */

package MaximumSubarray

//[-2,1,-3,4,-1,2,1,-5,4]
//[0, 3, -4, 7, -5, 3, -1, -6, 9]
func maxSubArray(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}
	dp := make([]int, l)

	dp[0] = nums[0]
	res := dp[0]

	for i := 1; i < l; i++ {
		if dp[i - 1] > 0 {
			dp[i] = dp[i - 1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
