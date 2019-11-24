/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-23 23:30
 */

package BaseballGame

import "strconv"

func calPoints(ops []string) int {
	res := 0
	stack := []int{}
	for i := 0; i < len(ops); i++ {
		if ops[i] == "C" {
			stack = stack[:len(stack) - 1]
		} else if ops[i] == "D" {
			stack = append(stack, 2 * stack[len(stack) - 1])
		} else if ops[i] == "+" {
			stack = append(stack, stack[len(stack) - 1] + stack[len(stack) - 2])
		} else {
			v, _ := strconv.Atoi(ops[i])
			stack = append(stack, v)
		}
	}

	for len(stack) != 0 {
		res += stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
	}
	return res
}