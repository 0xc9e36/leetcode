/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-13 21:56
 */

package MajorityElement


//func majorityElement(nums []int) int {
//
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	m := make(map[int]int, 0)
//	for i := 0; i < len(nums); i++ {
//		if c, ok := m[nums[i]] ; ok {
//			m[nums[i]] = c + 1
//			if c + 1 > len(nums) / 2 {
//				return nums[i]
//			}
//		} else {
//			m[nums[i]] = 1
//		}
//	}
//	return 1
//}



//摩尔投票
func majorityElement(nums []int) int {
	ans, count:= 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[ans] {
			count--
			if count == 0 {
				ans = i
				count = 1
			}
		} else {
			count++
		}
	}
	return nums[ans]
}