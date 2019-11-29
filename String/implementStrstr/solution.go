/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-27 21:02
 */



package implementStrstr



//流氓解法
//func strStr(haystack string, needle string) int {
//
//	n, m := len(haystack), len(needle)
//	for i := 0 ; i < n - m + 1; i++ {
//		j := 0
//		for ; j < m; j++ {
//			if haystack[i + j] != needle[j] {
//				break
//			}
//		}
//		if j == m {
//			return i
//		}
//	}
//	return -1
//}


// 1 2 3 4
// 1 2 4 3



// kmp 匹配
//func nextTable(p string) []int {
//	l := len(p)
//	if l == 0 {
//		return nil
//	}
//	next := make([]int, l, l)
//	next[0] = -1
//	k, j := -1, 0
//	for j < l - 1 {
//		if k == -1 || p[k] == p[j] {
//			k++
//			j++
//			next[j] = k
//		} else {
//			k = next[k]
//		}
//	}
//	return next
//}
//
//func strStr(haystack string, needle string) int {
//	m, n := len(haystack),len(needle)
//	i, j := 0, 0
//	next := nextTable(needle)
//	for i < m && j < n {
//		if j == -1 || haystack[i] == needle[j] {
//			i++
//			j++
//		} else {
//			j = next[j]
//		}
//	}
//	if j == n {
//		return i - j
//	}
//	return -1
//}