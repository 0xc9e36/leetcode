/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-21 23:47
 */

package LargestNumberAtLeastTwiceofOthers


func dominantIndex(nums []int) int {
	l := len(nums)
	m := make([]int, 101)

	//最大索引
	max := 0
	for i := 0; i < l; i++ {
		m[nums[i]]++
		if nums[i] > nums[max] {
			max = i
		}
	}

	for i := 1; i < 101; i++ {
		if m[i] == 0 {
			continue
		}

		if i != nums[max] && nums[max] / i < 2 {
			return -1
		}
	}

	return max
}
