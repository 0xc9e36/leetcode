/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-10 21:21
 */

package MinimumTimeVisitingAllPoints


func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for i := 0; i < len(points) - 1; i++ {
		dx := abs(points[i][0] - points[i + 1][0])
		dy := abs(points[i][1] - points[i + 1][1])
		ans += max(dx, dy)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}