/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-11 08:02
 */

package MaximumProductofThreeNumbers

import "sort"

func maximumProduct(nums []int) int {
	l := len(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return max(nums[0] * nums[1] * nums[l - 1], nums[l - 1] * nums[l - 2] * nums[l - 3])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}


