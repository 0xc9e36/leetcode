/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-10 22:17
 */

package HeightChecker


//暴力解法
//func heightChecker(heights []int) int {
//
//	bat := make([]int, len(heights))
//	copy(bat, heights)
//	sort.Slice(bat, func(i, j int) bool {
//		return bat[i] < bat[j]
//	})
//
//	ans := 0
//	for i := 0; i < len(bat); i++ {
//		if bat[i] !=  heights[i] {
//			ans++
//		}
//	}
//
//	return ans
//}


func heightChecker(heights []int) int {
	l, bucket := len(heights), [101]int{}
	for i := 0; i < l; i++ {
		bucket[heights[i]]++
	}

	j, ans := 0, 0
	for i := 1; i <= 100; i++ {
		for bucket[i] > 0 {
			if heights[j] != i {
				ans++
			}
			j++
			bucket[i]--
		}
	}
	return ans
}

