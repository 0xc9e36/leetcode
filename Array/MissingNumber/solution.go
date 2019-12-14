/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-12 20:46
 */

package MissingNumber


//暴力法
//func missingNumber(nums []int) int {
//	l := len(nums)
//	sum := l * (l + 1) / 2
//	for i := 0; i < l; i++ {
//		sum -= nums[i]
//	}
//	return sum
//}
//
//



func missingNumber(nums []int) int {
	l := len(nums)
	ans := 0 ^ nums[0]
	for i := 1; i < l; i++ {
		ans ^= i ^ nums[i]
	}
	return ans
}






