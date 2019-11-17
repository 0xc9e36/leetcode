/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-20 22:37
 */

package CanPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {

	if n <= 0 {
		return true
	}

	i, l := 0, len(flowerbed)


	if l == 1 {
		if n == 1 && flowerbed[0] == 0 {
			return true
		}
		return false
	}

	for i < l {
		if n <= 0 {
			return true
		}
		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else if i == l - 1 {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}
		i++
	}
	return n <= 0
}
