/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-21 22:56
 */

package MaximumAverageSubarrayI

///1,12,-5,-6,50,3

///11 -17 -1 56 -47
//笨办法
func findMaxAverage(nums []int, k int) float64 {
	l := len(nums)
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	res := sum
	for i := k; i < l; i++ {
		sum = sum + nums[i] - nums[i-k]
		res = max(res, sum)
	}

	return float64(res) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
