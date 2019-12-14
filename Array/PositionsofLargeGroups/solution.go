/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-12 20:26
 */

package PositionsofLargeGroups


func largeGroupPositions(S string) [][]int {
	ans := make([][]int, 0)
	l := len(S)
	for i := 0; i < l; i++ {
		j := i
		for i < l - 1 && S[i + 1] == S[i] {
			i++
		}
		if i - j >= 2 {
			ans = append(ans, []int{j, i})
		}
	}
	return ans
}