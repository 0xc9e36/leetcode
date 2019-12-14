/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-13 13:20
 */

package GasStation


//func canCompleteCircuit(gas []int, cost []int) int {
//	l := len(gas)
//	for i := 0; i < l; i++ {
//		curr := i
//		sum := gas[i]
//		for sum >= cost[curr] {
//			sum -= cost[curr]
//			curr = (curr + 1) % l
//			if curr == i {
//				return i
//			}
//			sum += gas[curr]
//		}
//	}
//	return -1
//}


func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	var total, sum int
	ans := 0
	for i := 0; i < l; i++ {
		sum += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if sum < 0 {
			sum = 0
			ans = i + 1
		}
	}

	if total < 0 {
		return -1
	}
	return ans
}