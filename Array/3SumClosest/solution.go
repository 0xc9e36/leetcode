package _SumClosest

import "sort"

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	l, result := len(nums), nums[0] + nums[1] + nums[2]
	for i := 0; i < l - 2; i++ {
		lo, hi := i + 1, l - 1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			if sum < target {
				lo++
			} else if sum > target {
				hi--
			} else {
				return target
			}

			if abs(sum - target) < abs(result - target) {
				result = sum
			}
		}
	}
	return result
}


//个人觉得这个语言有点2，这么基础的函数居然都要自己实现
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}