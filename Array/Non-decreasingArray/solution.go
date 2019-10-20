/**
 * @Author: wei.tan
 * @Description:
 * @File:  solutio
 * @Version: 1.0.0
 * @Date: 2019-10-20 13:20
 */

package Non_decreasingArray

//1 3 2
//4 2 3
//2 4 3
//4 4 3

//三种情况:
//1. n == 0, nums[n] > nums[n + 1], 修改 nums[n] 为 nums[n + 1]
//2. n != 0, nums[n] > nums[n + 1]
//	2.1 nums[n + 1] >= nums[n - 1], 修改 nums[n] 为 nums[n + 1]
//	2.2 nums[n + 1] < nums[n - 1], 修改 nums[n + 1] 为 nums[n]

func checkPossibility(nums []int) bool {
	l := len(nums)
	if l < 2 {
		return true
	}

	modified := false

	for i := 0; i < l-1; i++ {
		//出现逆序
		if nums[i] > nums[i+1] {
			//已经修改过
			if modified {
				return false
			} else {
				if i == 0 {
					nums[i] = nums[i+1]
				} else if nums[i+1] > nums[i-1] {
					nums[i] = nums[i+1]
				} else {
					nums[i+1] = nums[i]
				}
				modified = true
			}
		}
	}
	return true
}
