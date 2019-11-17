/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-22 20:28
 */

package FindPivotIndex

// [1, 7, 3, 6, 5, 6]

func pivotIndex(nums []int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	left, right := make([]int, l), make([]int, l)
	ls, rs := 0, 0
	for i := 1; i < l; i++ {
		ls += nums[i-1]
		left[i] = ls
		rs += nums[l-i]
		right[l-1-i] = rs
	}

	for i := 0; i < l; i++ {
		if left[i] == right[i] {
			return i
		}
	}

	return -1
}
