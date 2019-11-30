/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-30 12:19
 */

package GroupAnagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	tab := make(map[string][]string, 0)
	for i := 0; i < len(strs); i++ {
		w := stringSort(strs[i])
		tab[w] = append(tab[w], strs[i])
	}

	res := [][]string{}
	for _, indexs := range tab {
		res = append(res, indexs)
	}
	return res
}

func stringSort(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
