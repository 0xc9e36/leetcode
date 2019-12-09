/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-08 22:41
 */

package RemoveDuplicatesfromSortedArrayII


func removeDuplicates(nums []int) int {
	k, l := 0, len(nums)
	if l == 0 {
		return 0
	}

	for i := 0; i <= l - 1; i++ {
		j := i
		for i < l - 1 && nums[i] == nums[i + 1] {
			i++
		}

		if i == j {
			nums[k] = nums[i]
			k++
		} else {
			nums[k] = nums[j]
			nums[k + 1] = nums[j + 1]
			k += 2
		}
	}
	return k
}
