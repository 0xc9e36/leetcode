/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-12 21:32
 */

package TwoSumII_Inputarrayissorted



func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers) - 1
	ans := []int{}
	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			ans = append(ans, start + 1)
			ans = append(ans, end + 1)
			break
		} else if sum < target {
			start++
		} else {
			end--
		}
	}
	return ans
}