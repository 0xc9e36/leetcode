/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-22 22:10
 */

package LongestContinuousIncreasingSubsequence

// [1,3,5,4,7]
// 1 1 2
func findLengthOfLCIS(nums []int) int {
	l, i, max := len(nums), 1, 1
	if l == 0 {return 0}
	for i < l {
		j := i
		for i < l && nums[i] > nums[i - 1] {
			i++
		}

		if i - j + 1 > max {
			max = i - j + 1
		}

		i++
	}

	return max
}
