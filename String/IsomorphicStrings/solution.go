/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-30 12:02
 */

package IsomorphicStrings

//我最开始写的
//func isIsomorphic(s string, t string) bool {
//	tab := [256]uint16{}
//	m := map[byte]byte{}
//	n := map[byte]byte{}
//	for i := 0; i < len(s); i++ {
//		c, ok := m[s[i]]
//		if !ok {
//			m[s[i]] = t[i]
//		} else if c != t[i] {
//			return false
//		}
//
//		ch, ok2 := n[t[i]]
//		if !ok2 {
//			n[t[i]] = s[i]
//		} else if ch != s[i] {
//			return false
//		}
//		tab[t[i]]++
//	}
//
//	for i := 0; i < len(s); i++ {
//		c, ok := m[s[i]]
//		if !ok {
//			return false
//		}
//		tab[c]--
//	}
//
//	for i := 0; i < 256; i++ {
//		if tab[i] != 0 {
//			return false
//		}
//	}
//	return true
//}


//大神们简洁的代码
func isIsomorphic(s string, t string) bool {
	ms, mt := [256]int{}, [256]int{}
	for i := 0; i < len(s); i++ {
		if ms[s[i]] != mt[t[i]] {
			return false
		}
		ms[s[i]] = i + 1
		mt[t[i]] = i + 1
	}
	return true
}