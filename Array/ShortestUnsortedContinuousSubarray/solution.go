/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-20 21:47
 */

package ShortestUnsortedContinuousSubarray

import "sort"

// 5 2 1 4

//1 7 6 8 2
//[2,6,4,8,10,9,15]

//解法一、直接排序对比
func findUnsortedSubarray(nums []int) int {
	cp := make([]int, len(nums))
	copy(cp, nums)
	sort.Ints(cp)

	var start, end int
	l := len(nums)
	for i := 0; i < l; i++ {
		if cp[i] != nums[i] {
			start = i
			break
		}
	}

	for i := l - 1; i >= 0; i-- {
		if cp[i] != nums[i] {
			end = i
			break
		}
	}

	// 1 3 2
	if start >= end {
		return 0
	}
	return end - start + 1
}