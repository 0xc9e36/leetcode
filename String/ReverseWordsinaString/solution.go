/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-29 23:00
 */

package ReverseWordsinaString

import "strings"


//翻转两次
func reverseWords(s string) string {

	b := []byte(s)
	reverse(b)

	for i := 0; i < len(b); i++ {
		j := i
		for i < len(b) && b[i] == 32 {
			i++
		}
		if i == j || i - j == 1 {
			continue
		}
		b =  append(b[:j + 1], b[i:]...)
		i--
	}

	for i := 0; i < len(b); i++ {
		j := i
		for i < len(b) && b[i] != 32 {
			i++
		}
		if j != i {
			reverse(b[j:i])
		}
	}
	return strings.TrimSpace(string(b))
}

func reverse(b []byte) {
	i, j := 0, len(b) - 1
	for i < j {
		b[i], b[j]  = b[j], b[i]
		i++
		j--
	}
}