package _Sum

import (
	"sort"
)

//固定开始一位
func threeSum(nums []int) [][]int {
	//排序
	sort.Ints(nums)

	l, result := len(nums), make([][]int, 0)

	for i := 0; i < l-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo, hi := i+1, l-1
			for lo < hi {
				if nums[i]+nums[lo]+nums[hi] < 0 {
					lo++
				} else if nums[i]+nums[lo]+nums[hi] > 0 {
					hi--
				} else {
					result = append(result, []int{nums[i], nums[lo], nums[hi]})

					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					lo++
					hi--
				}
			}
		}

	}

	return result
}
