/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-05 10:52
 */

package IsSubsequence


func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] != t[j] {
			j++
			continue
		}
		i++
		j++
	}
	if i == len(s) {
		return true
	}
	return false
}