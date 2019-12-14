/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-12 20:11
 */

package BinaryPrefixDivisibleBy5


func prefixesDivBy5(A []int) []bool {
	l := len(A)
	ans := make([]bool, l)
	product := A[0]
	ans[0] = product % 5 == 0
	for i := 1; i < l; i++ {
		product = product<<1 + A[i]
		//能被 5 整除的不用关心
		product %= 5
		ans[i] = product % 5 == 0

	}
	return ans
}
