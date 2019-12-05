/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-05 09:13
 */

package BestTimetoBuyandSellStock


func maxProfit(prices []int) int {
	l := len(prices)
	res := 0
	if l == 0 {
		return 0
	}
	dp := make([]int, l, l)
	dp[0] = prices[0]
	for i := 1; i < l; i++ {

		if prices[i] - dp[i - 1] > res {
			res = prices[i] - dp[i - 1]
		}

		if prices[i] < dp[i - 1] {
			dp[i] = prices[i]
		} else {
			dp[i] = dp[i - 1]
		}
	}
	return res
}