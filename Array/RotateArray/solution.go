/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-20 23:09
 */

package RotateArray


//解法一
//func rotate(nums []int, k int)  {
//	l := len(nums)
//	for i := 0; i < k; i++ {
//		t := nums[l - 1]
//		for j := l - 1; j >= 1; j-- {
//			nums[j] = nums[j - 1]
//		}
//		nums[0] = t
//	}
//}


//解法二
//[1,2,3,4,5,6,7]  k = 3
//[5,6,7,1,2,3,4]

//7 6 5 4 3 2 1  ==> 5 6 7 1 2 3 4

// n => k - 1
// n - 1 => k - 2
// n - k + 1 => 0


func rotate(nums []int, k int)  {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int)  {
	start, end := 0, len(nums) - 1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
