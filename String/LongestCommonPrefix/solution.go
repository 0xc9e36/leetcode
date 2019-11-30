/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-29 22:17
 */

package LongestCommonPrefix


func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}
	common := strs[0]


	for i := 0; i < len(common); i++ {
		j := 1
		for j < len(strs) {
			if i >= len(strs[j]) {
				break
			}
			if strs[j][i] != common[i] {
				break;
			}
			j++
		}
		if j == len(strs) {
			res += string(common[i])
		} else {
			break
		}
	}
	return res
}