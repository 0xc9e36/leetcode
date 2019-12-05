/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-05 13:42
 */

package CountingBits

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	dp := make([]int, num + 1, num + 1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= num; i++ {
		if i % 2 == 0 {
			dp[i] = dp[i>>1]
		} else {
			dp[i] = dp[i>>1] + 1
		}
	}
	return dp
}
