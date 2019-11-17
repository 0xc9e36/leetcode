/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-22 20:50
 */

package AddtoArray_FormofInteger

import "strconv"

//[1, 2, 0, 0]
//		[3, 4]

// 4 3
// 0 0  1 2
func addToArrayForm(A []int, K int) []int {

	//转为数组
	ks := make([]int, len(strconv.Itoa(K)))
	i := 0
	for K > 0 {
		b := K % 10
		ks[i] = b
		K /= 10
		i++
	}

	reverse(A)

	la, ls := len(A), len(ks)
	res, f := []int{}, 0
	for i := 0; i < la || i < ls; i++ {
		if i < la {
			f += A[i]
		}

		if i < ls {

			f += ks[i]
		}

		res = append(res, f%10)
		f /= 10
	}

	if f != 0 {
		res = append(res, f)
	}

	reverse(res)

	return res
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

