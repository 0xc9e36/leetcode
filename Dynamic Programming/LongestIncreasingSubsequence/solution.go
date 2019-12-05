/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-03 20:28
 */

package LongestIncreasingSubsequence

func divisorGame(N int) bool {
	if N == 1 {
		return false
	}

	//dp[i] 表示面对 number 为 i, 是否先手必胜
	dp := make([]bool, N + 1, N + 1)
	dp[1] = false

	for i := 2; i <= N; i++ {
		dp[i] = false
		for x := 1; x < i; x ++ {
			if i % x == 0 && !dp[i - x] {
				dp[i] = true
				break
			}
		}
	}
	return dp[N]
}
