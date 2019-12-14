/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-12 22:25
 */

package ToeplitzMatrix


func isToeplitzMatrix(matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n - 1; i++{
		for j := 0; j < m - 1; j++ {
			if matrix[i][j] != matrix[i + 1][j + 1] {
				return false
			}
		}
	}
	return true
}