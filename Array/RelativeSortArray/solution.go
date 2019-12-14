/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-10 23:09
 */

package RelativeSortArray


func relativeSortArray(arr1 []int, arr2 []int) []int {

	l1, l2 := len(arr1), len(arr2)
	tab := [1001]int{}
	for i := 0; i < l1; i++ {
		tab[arr1[i]]++
	}

	j := 0
	for i := 0; i < l2; i++ {
		l := tab[arr2[i]]
		for k := 1; k <= l; k++ {
			arr1[j] = arr2[i]
			j++
			tab[arr2[i]]--
		}
	}

	for i := 0; i <= 1000; i++ {
		for k := 1; k <= tab[i]; k++ {
			arr1[j] = i
			j++
		}
	}

	return arr1
}

