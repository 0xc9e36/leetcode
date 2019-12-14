/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-13 09:13
 */

package FirstMissingPositive



func firstMissingPositive(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 1
	}

	// 0 1 2  3
	// 3 4 -1 1
	for i := 0; i < l; i++ {
		for nums[i] > 0 && nums[i] <= l && nums[i] != nums[nums[i] - 1] {
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		}
	}

	ans := 0
	for ; ans < l; ans++ {
		if nums[ans] != ans + 1 {
			break
		}
	}

	return ans + 1
}