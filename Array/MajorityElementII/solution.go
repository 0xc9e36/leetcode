/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-13 23:29
 */

package MajorityElementII


func majorityElement(nums []int) []int {
	n1, n2, c1, c2 := 0, 0, 0, 0

	for i := 0 ; i < len(nums); i++ {
		if nums[i] == n1 {
			c1++
		} else if nums[i] == n2 {
			c2++
		} else if c1 == 0 {
			n1 = nums[i]
			c1 = 1
		} else if  c2 == 0 {
			n2 = nums[i]
			c2 = 1
		} else {
			c1--
			c2--
		}
	}

	c1, c2 = 0, 0
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == n1 {
			c1++
		}

		if nums[i] == n2 {
			c2++
		}
	}

	if c1 > len(nums) / 3 {
		ans = append(ans, n1)
	}

	if n1 != n2 && c2 > len(nums) / 3 {
		ans = append(ans, n2)
	}
	return ans
}

// 1 2 1 2 1