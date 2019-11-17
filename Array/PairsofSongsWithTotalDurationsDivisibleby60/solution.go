/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-23 08:51
 */

package PairsofSongsWithTotalDurationsDivisibleby60

func numPairsDivisibleBy60(time []int) int {
	l := len(time)
	if l == 1 {
		return 0
	}

	res := 0
	m := map[int]int{}
	for i := 0; i < l; i++ {
		m[time[i]] = i
	}

	for i := 0; i < l; i++ {
		if i == 0 {
			continue
		}

		t := time[i] % 60

	}
	return res
}