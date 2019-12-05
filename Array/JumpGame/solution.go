/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-11-30 19:17
 */

package JumpGame

func canJump(nums []int) bool {
	l := len(nums)
	if l == 0 {
		return true
	}
	f := make([]bool, l, l)
	f[0] = true
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			f[i] = f[j] && nums[j] >= i - j
			if f[i] {
				break
			}
		}
	}


	return f[l - 1]
}
