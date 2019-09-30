package JewelsandStones

func numJewelsInStones(J string, S string) int {
	m := [128]int{}
	for i := range J{
		m[J[i]]++
	}

	result := 0

	for j := range S {
		if m[S[j]] != 0 {
			result++
		}
	}
	return result
}
