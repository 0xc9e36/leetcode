/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-05 12:23
 */

package HouseRobber

func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	dp := make([]int, l, l)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
	}
	return dp[l - 1]
}


func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
