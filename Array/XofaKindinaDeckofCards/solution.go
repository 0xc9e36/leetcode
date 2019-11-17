/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-20 23:31
 */

package XofaKindinaDeckofCards


//[1,1,1,2,2,2,3,3]
func hasGroupsSizeX(deck []int) bool {

	bm := [1000]int{}

	for _, d := range deck {
		bm[d]++
	}

	x := 1000
	for i := 0; i < 1000; i++ {

		if bm[i] == 1 {
			return false
		}

		if bm[i] == 0 {
			continue
		}

		if bm[i] < x {
			x = bm[i]
		}
	}

	for i := 0; i < 1000; i++ {

		if bm[i] == 0 {
			continue
		}

		if bm[i] % x == 1 {
			return false
		}
	}
	return true
}


