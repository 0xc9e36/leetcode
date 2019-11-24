/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-24 09:46
 */

package ValidateStackSequences


func validateStackSequences(pushed []int, popped []int) bool {

	i := 0
	stack := []int{}
	for j := 0; j < len(popped); j++{
		for len(stack) == 0 || popped[j] != stack[len(stack) - 1] {
			if i == len(pushed) {
				return false
			}
			stack = append(stack, pushed[i])
			i++
		}
		stack = stack[:len(stack) - 1]
	}
	return true
}


