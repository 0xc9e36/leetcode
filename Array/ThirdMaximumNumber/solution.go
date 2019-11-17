/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-20 15:30
 */

package ThirdMaximumNumber

func thirdMax(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	first, second, third := nums[0], nums[0], nums[0]
	for i := 0; i < l; i++ {
		if nums[i] > first {
			third = second
			second = first
			first = nums[i]
		} else if nums[i] != first && nums[i] > second || second == first {
			third = second
			second = nums[i]
		} else if nums[i] != first && nums[i] != second && nums[i] > third || third == second || third == first {
			third = nums[i]
		}
	}

	if first > second && second > third {
		return third
	}

	return first
}