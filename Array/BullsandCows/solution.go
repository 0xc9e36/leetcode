/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-13 12:57
 */

package BullsandCows

import "fmt"

func getHint(secret string, guess string) string {
	l := len(secret)
	tab := [10]int{}
	for i := 0; i < l; i++ {
		tab[secret[i] - '0']++
	}

	common := 0
	for i := 0; i < l; i++ {
		if tab[guess[i] - '0'] > 0 {
			common++
			tab[guess[i] - '0']--
		}
	}

	A := 0
	for i := 0; i < l; i++ {
		if secret[i] == guess[i] {
			A++
		}
	}

	return fmt.Sprintf("%dA%dB", A, common - A)
}