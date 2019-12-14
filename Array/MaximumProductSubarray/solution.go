/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-11 13:07
 */

package MaximumProductSubarray

func maxProduct(nums []int) int {
	l := len(nums)
	curr_min, curr_max, ans := nums[0],  nums[0],  nums[0]
	for i:= 1; i < l; i++ {
		m1, m2 := curr_min * nums[i], curr_max * nums[i]
		curr_min = min(min(m1, m2), nums[i])
		curr_max = max(max(m1, m2), nums[i])
		ans = max(ans, curr_max)
	}

	return ans
}


func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}