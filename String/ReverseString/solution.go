/**
 * @Author: wei.tan
 * @Description:
 * @File:  string
 * @Version: 1.0.0
 * @Date: 2019-11-29 22:58
 */

package ReverseString

func reverseString(s []byte)  {
	i, j := 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}