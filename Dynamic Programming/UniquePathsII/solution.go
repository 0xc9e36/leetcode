/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-30 20:43
 */

package UniquePathsII


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else if i == 0 && j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
				if i >= 1 {
					dp[i][j] += dp[i - 1][j]
				}
				if j >= 1 {
					dp[i][j] += dp[i][j - 1]
				}
			}
		}
	}

	return dp[m - 1][n - 1]
}

