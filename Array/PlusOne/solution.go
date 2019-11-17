/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-22 20:10
 */

package PlusOne

// 9 9
func plusOne(digits []int) []int {
	res := []int{0}
	res = append(res, digits...)
	l := len(res)

	f := 1
	for i := l - 1; i >= 0; i-- {
		t := res[i] + f
		if t >= 10 {
			res[i] = (t) % 10
			f = t / 10
		} else {
			res[i] = res[i] + f
			f = 0
		}
	}
	if res[0] == 0 {
		res = res[1:]
	}
	return res
}
