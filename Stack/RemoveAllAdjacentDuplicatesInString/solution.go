/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-23 23:06
 */

package RemoveAllAdjacentDuplicatesInString

func removeDuplicates(S string) string {
	stack := []uint8{}
	for i := range S {
		c := S[i] - 97
		if len(stack) == 0 || stack[len(stack) - 1] != c {
			stack = append(stack, c)
		} else {
			stack = stack[:len(stack) - 1]
		}
	}

	res := ""
	for len(stack) != 0 {
		res = string(stack[len(stack) - 1] + 97) + res
		stack = stack[:len(stack) - 1]
	}
	return res

}