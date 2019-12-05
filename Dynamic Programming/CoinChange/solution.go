/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-30 15:23
 */

package CoinChange


func coinChange(coins []int, amount int) int {
	f := make([]int, amount + 1, amount + 1)
	f[0] = 0
	max := amount + 1
	for i := 1; i <= amount; i++ {
		//不能拼用 max 表示
		f[i] = max
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && f[i - coins[j]] != max && f[i - coins[j]] + 1 < f[i] {
				f[i] = f[i - coins[j]] + 1
			}
		}
	}
	if f[amount] == max {
		return -1
	}
	return f[amount]
}
