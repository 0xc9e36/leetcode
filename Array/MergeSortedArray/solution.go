/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-21 22:39
 */

package MergeSortedArray

//Input:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//Output: [1,2,2,3,5,6]

func merge(nums1 []int, m int, nums2 []int, n int) {

	k := len(nums1)
	l := m + n
	m, n = m - 1, n - 1

	for {
		k--
		if m >= 0 && n >= 0 {
			if nums1[m] >= nums2[n] {
				nums1[k] = nums1[m]
				m--
			} else {
				nums1[k] = nums2[n]
				n--
			}
		} else if m >= 0 {
			nums1[k] = nums1[m]
			m--
		} else if n >= 0 {
			nums1[k] = nums2[n]
			n--
		} else {
			nums1 = nums1[:l]
			break
		}
	}
}
