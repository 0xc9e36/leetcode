/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-23 08:27
 */

package NumberofEquivalentDominoPairs

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	l := len(dominoes)

	if l <= 1 {
		return 0
	}

	for i := 0; i < l; i++ {
		if dominoes[i][0] > dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}
	}

	m := map[string]int{}

	res := 0
	for i := 0; i < l; i++ {
		k := fmt.Sprintf("%d:%d", dominoes[i][0], dominoes[i][1])
		j, ok := m[k]
		if !ok {
			m[k] = 1
		} else {
			res += j
			m[k] = j + 1
		}
	}
	return res
}
