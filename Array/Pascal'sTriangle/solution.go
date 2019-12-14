/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-12 20:38
 */

package Pascal_sTriangle


func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, numRows)
	}

	for i := 0; i < numRows; i++ {
		ans[i][0] = 1
	}
	for i := 1; i < numRows; i++ {
		for j := 1; j < numRows; j++ {
			ans[i][j] = ans[i - 1][j - 1] + ans[i - 1][j]
		}
	}

	for i := 0; i < numRows; i++ {
		ans[i] = ans[i][:i + 1]
	}
	return ans
}


// 1 0 0 0 0
// 1 1 0 0 0
// 1 2 1 0 0
// 1 3 3 1 0
// 1 4 6 4 1