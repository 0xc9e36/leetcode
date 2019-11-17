/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-21 22:11
 */

package ContainsDuplicateII

//1 2 1 2 3 3 4 5

// k 滑动窗口

// k = 2
// 2 3
func containsNearbyDuplicate(nums []int, k int) bool {

	l := len(nums)

	if k == 0 {
		return false
	} else if k >= l {
		k = l
	}

	wd := make([]int, k)
	m := make(map[int]struct{})
	for i := 0; i < k; i++ {
		wd[i] = nums[i]
		_, ok := m[nums[i]]
		if ok {
			return true

		}
		m[nums[i]] = struct{}{}
	}

	//1 2 3   k =
	for j := k; j < l; j++ {
		_, ok := m[nums[j]]
		if ok {
			return true
		}

		delete(m, wd[0])
		wd = append(wd, nums[j])
		wd = wd[1:]
		m[nums[j]] = struct{}{}
	}

	return false
}

//暴力
//func containsNearbyDuplicate(nums []int, k int) bool {
//	l := len(nums)
//	for i := 0; i < l; i++ {
//
//		//左边
//		for j := 1; j <= k; j++ {
//			if i - j < 0 {
//				break
//			}
//			if nums[i] == nums[i - j] {
//				return true
//			}
//		}
//
//		//右边
//		for j := 1; j <= k; j++ {
//			if i + j >= l {
//				break
//			}
//			if nums[i] == nums[i + j] {
//				return true
//			}
//		}
//	}
//
//	return false
//}
