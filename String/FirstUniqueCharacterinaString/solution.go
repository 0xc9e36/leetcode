/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-29 22:45
 */

package FirstUniqueCharacterinaString

//没什么难度
func firstUniqChar(s string) int {
	tab := [26]uint16{}
	for i := 0; i < len(s); i++ {
		tab[s[i] - 97]++
	}
	for j := 0; j < len(s); j++ {
		if tab[s[j] - 97] == 1 {
			return j
		}
	}
	return -1
}