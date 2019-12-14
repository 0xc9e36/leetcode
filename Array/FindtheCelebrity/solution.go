/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-09 21:53
 */

package FindtheCelebrity


/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		ans := 0
		for i := 1; i < n; i++ {
			if knows(ans, i) {
				ans = i
			}
		}

		for i := 0; i < n; i++ {
			if i != ans {
				if knows(i, ans) && !knows(ans, i) {

				} else {
					return -1
				}
			}
		}
		return ans
	}
}