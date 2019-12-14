/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-12-12 20:38
 */

package __bitand2_bitCharacters

func isOneBitCharacter(bits []int) bool {
	l := len(bits)
	for i := 0; i < l; i++ {
		if bits[i] == 0 {
			if i == l - 1 {
				return true
			}
		} else {
			if i == l - 1 || i == l - 2 {
				return false
			} else {
				i++
			}
		}
	}
	return false
}