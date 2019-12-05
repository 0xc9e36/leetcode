/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-05 10:34
 */

package MinCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	dp := make([]int, l + 1, l + 1)
	for i := 2; i <= l; i++ {
		dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
	}
	return dp[l]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}