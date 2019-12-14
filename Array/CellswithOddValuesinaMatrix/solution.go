/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-10 21:50
 */

package CellswithOddValuesinaMatrix

func oddCells(n int, m int, indices [][]int) int {
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
	}

	for i := 0; i < len(indices); i++ {
		x, y := indices[i][0], indices[i][1]
		A[x][y]++
		for j := 0; j < m; j++ {
			A[x][j]++
		}
		for k := 0; k < n; k++ {
			if k == x {
				continue
			}
			A[k][y]++
		}

	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] % 2 == 1 {
				ans++
			}
		}
	}
	return ans
}