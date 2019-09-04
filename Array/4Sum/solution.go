package _Sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	l, result := len(nums), [][]int{}

	// sort
	sort.Ints(nums)

	//先排序，再固定两个数, 然后头尾指针扫描剩下的元素
	for i := 0; i < l - 3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < l - 2; j++ {
			if j != i + 1 && nums[j] == nums[j-1] {
				continue
			}
			lo, hi := j + 1, l - 1
			for lo < hi {
				sum := nums[i] + nums[j] + nums[lo] + nums[hi]
				if sum == target {
					result = append(result, []int{nums[i],nums[j],nums[lo],nums[hi]})
					lo++
					hi--
					for lo < hi && nums[lo] == nums[lo - 1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi + 1] {
						hi--
					}
				} else if sum < target {
					lo++
					for lo < hi && nums[lo] == nums[lo - 1] {
						lo++
					}
				} else {
					hi--
					for lo < hi && nums[hi] == nums[hi + 1] {
						hi--
					}
				}
			}
		}
	}
	return result
}