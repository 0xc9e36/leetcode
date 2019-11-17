/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-20 18:01
 */

package MaximizeDistancetoClosestPerson


//[1,0,0,0,0,1,0,1]

//1 0 0 1
//[0,1,0,0,0,0,0,0,1,1,0,1,1]
//[1,0,0,0,0,1,0,1]
func maxDistToClosest(seats []int) int {

	i, l, res := 0, len(seats), 1

	for i < l {
		j := i

		for j < l && seats[j] == 0 {
			j++
		}

		if i == 0 || j == l {
			res = max(res, j - i)
		} else if j - i > res {
			var v int
			if (j - i) & 1 == 1 {
				v = ((j - i) / 2) + 1
			} else {
				v = (j - i) / 2
			}
			res = max(res, v)
		}

		if j == i {
			i++
		} else {
			i = j + 1
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}