/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-23 21:06
 */

package DegreeofanArray

import "math"

func findShortestSubArray(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return len(nums)
	}
	tab, degree := [50001]int{}, 0
	for i := 0; i < len(nums); i++ {
		tab[nums[i]]++
		if tab[nums[i]] > degree {
			degree = tab[nums[i]]
		}
	}
	ls := []int{}
	for i := 0; i <= 50000; i++ {
		if tab[i] == degree {
			ls = append(ls, i)
		}
	}

	// 1 2 2 1
	ans := math.MaxInt32
	for i := range ls {
		start, end := 0, len(nums) - 1
		ans = min(ans, end - start + 1)
		for start < end {
			if nums[start] == ls[i] && nums[end] == ls[i] {
				break
			} else if nums[start] != ls[i] {
				start++
			} else {
				end--
			}
			ans = min(ans, end - start + 1)
		}
	}
	return ans
}



func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
