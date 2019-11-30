/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-29 22:43
 */

package LengthofLastWord

func lengthOfLastWord(s string) int {
	l := len(s)

	for j := l - 1; j >= 0; j-- {
		if s[j] != ' ' {
			end := j
			for j >= 0 && s[j] != ' ' {
				j--
			}
			return end - j
		}
	}

	return 0
}