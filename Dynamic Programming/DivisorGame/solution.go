/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-30 20:44
 */

package DivisorGame

import "strconv"

func numDecodings(s string) int {

	l := len(s)
	dp := make([]int, l + 1, l + 1)
	dp[0] = 1
	for i := 1; i <= l; i++ {
		dp[i] = 0
		if s[i - 1] >= '1' && s[i - 1] <= '9' {
			dp[i] += dp[i - 1]
		}
		if i > 1{
			t, _ := strconv.Atoi(s[i - 2: i])
			if t >= 10 && t <= 26 {
				dp[i] += dp[i - 2]
			}
		}
	}
	return dp[l]
}