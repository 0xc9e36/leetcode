/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-22 21:31
 */

package MaximumSubarray

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	//滚动数组
	dp := make([]int, 2)
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < l; i++ {
		dp[i%2] = max(dp[(i - 1) % 2] + nums[i], nums[i])
		ans = max(dp[i % 2], ans)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}