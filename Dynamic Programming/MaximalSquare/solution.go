/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-05 16:11
 */

package MaximalSquare


func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == '0' {
					dp[i][j] = 0
				} else {
					dp[i][j] = 1
				}
			} else if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(dp[i][j - 1], dp[i - 1][j])
				dp[i][j] = min(dp[i][j], dp[i - 1][j - 1]) + 1
			}

			if dp[i][j] * dp[i][j] > ans {
				ans = dp[i][j] * dp[i][j]
			}


		}
	}

	return  ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
