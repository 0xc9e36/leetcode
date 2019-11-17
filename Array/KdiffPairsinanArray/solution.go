/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-20 18:50
 */

package KdiffPairsinanArray

//[3, 1, 4, 1, 5], k = 2   3 1 4 5
//[1,3]  [3,5]

//[1,3,1,5,4]
//0
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	m := map[int]int{}
	for _, num := range nums {
		c, ok := m[num]
		if !ok {
			m[num] = 1
		} else {
			m[num] = c + 1
		}
	}

	res := 0
	for key, val := range m {
		if k == 0 && val > 1 {
			res++
		} else if k != 0 {
			_, ok := m[key + k]
			if ok {
				res++
			}
		}

	}
	return res
}
