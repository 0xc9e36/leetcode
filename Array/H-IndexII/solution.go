/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-14 09:34
 */

package H_IndexII


func hIndex(citations []int) int {
	start, end, l := 0, len(citations) - 1, len(citations)
	ans := 0
	for start <= end {
		mid := start + (end - start) >> 1
		if t := l - mid; citations[mid] >= t {
			ans = t
			end = mid -  1
		} else {
			start = mid + 1
		}
	}
	return ans
}
