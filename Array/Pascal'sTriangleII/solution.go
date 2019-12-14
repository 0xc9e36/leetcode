/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-22 22:30
 */

package Pascal_sTriangleII


//func getRow(rowIndex int) []int {
//	A := make([][]int, rowIndex + 1)
//	for i := 0; i <= rowIndex ; i++ {
//		A[i] = make([]int, rowIndex + 1)
//	}
//
//	for i := 0; i <= rowIndex; i++ {
//		A[i][0] = 1
//	}
//
//	for i := 1; i <= rowIndex; i++ {
//		for j := 1; j <= rowIndex; j++ {
//			A[i][j] = A[i - 1][j - 1] + A[i - 1][j]
//		}
//	}
//	return A[rowIndex]
//}



//滚动数组
func getRow(rowIndex int) []int {
	A := make([]int, rowIndex + 1)
	A[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			A[j] += A[j - 1]
		}
	}
	return A
}
