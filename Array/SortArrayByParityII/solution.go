/**
 * @Author: wei.tan
 * @Description:
 * @File:  solutioin
 * @Version: 1.0.0
 * @Date: 2019-12-10 22:07
 */

package SortArrayByParityII


//暴力解法
//func sortArrayByParityII(A []int) []int {
//	even, odd := 1, 0
//	l := len(A)
//	ans := make([]int, l)
//	for i := 0; i < l; i++ {
//		if A[i] & 1 == 0 {
//			ans[odd] = A[i]
//			odd += 2
//		} else {
//			ans[even] = A[i]
//			even += 2
//		}
//	}
//	return ans
//}
//


func sortArrayByParityII(A []int) []int {
	even, odd := 0, 1
	l := len(A)
	for i := 0; i < l; i++ {
		for odd < l && A[odd] & 1 == 1 {
			odd += 2
		}

		for even < l && A[even] & 1 == 0 {
			even += 2
		}
		if odd < l && even < l {
			A[odd], A[even] =  A[even], A[odd]
		}
	}
	return A
}
