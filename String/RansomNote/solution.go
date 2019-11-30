/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-29 22:52
 */

package RansomNote

func canConstruct(ransomNote string, magazine string) bool {
	tab := [26]int{}
	for i := 0; i < len(magazine); i++ {
		tab[magazine[i] - 97]++
	}
	for i := 0; i < len(ransomNote); i++ {
		tab[ransomNote[i] - 97]--
		if tab[ransomNote[i] - 97] < 0 {
			return false
		}
	}
	return true
}