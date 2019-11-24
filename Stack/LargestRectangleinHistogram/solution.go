/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-24 21:50
 */

package LargestRectangleinHistogram

func largestRectangleArea(heights []int) int {
	res := 0
	stack := make([]int, 0)
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		for len(stack) != 0 && heights[i] < heights[stack[len(stack) - 1]] {
			j := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			var area int
			if len(stack) == 0 {
				area = heights[j] * i
			} else {
				area = heights[j] * (i - stack[len(stack) - 1] - 1)
			}

			res = max(res, area)
		}
		stack = append(stack, i)
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}