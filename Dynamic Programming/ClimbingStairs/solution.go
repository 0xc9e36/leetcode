/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-05 11:32
 */

package ClimbingStairs

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n + 1, n + 1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}