/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-21 20:49
 */

package ValidMountainArray


// 1 2 3 4 5
// 1 2 5 4 3
//2 0 2
//[9,8,7,6,5,4,3,2,1,0]
func validMountainArray(A []int) bool {

	l := len(A)

	if l < 3 {
		return false
	}

	i := 0
	for i < l - 1 && A[i] < A[i + 1] {
		i++
	}

	if i == 0 || i == l - 1{
		return false
	}

	for i < l - 1 && A[i] > A[i + 1] {
		i++
	}

	if i != l - 1 {
		return false
	}

	return A[i] < A[i - 1]
}
