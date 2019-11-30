/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-30 11:19
 */

package ReverseVowelsofaString

func reverseVowels(s string) string {
	b := []byte(s)
	i, j := 0, len(b) - 1
	for i < j {
		for i < j && !isVowel(b[i]) {
			i++
		}
		for i < j && !isVowel(b[j]) {
			j--
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}


var m = map[byte]struct{}{
	'a' : struct{}{},
	'e' : struct{}{},
	'i' : struct{}{},
	'o' : struct{}{},
	'u' : struct{}{},
	'A' : struct{}{},
	'E' : struct{}{},
	'I' : struct{}{},
	'O' : struct{}{},
	'U' : struct{}{},
}

func isVowel(b byte) bool {
	_, ok := m[b]
	return ok
}