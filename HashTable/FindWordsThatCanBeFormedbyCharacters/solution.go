package FindWordsThatCanBeFormedbyCharacters

func countCharacters(words []string, chars string) int {
	res := 0
	for i := range words {
		m := [26]int{}
		for j := range chars {
			m[chars[j] - 97]++
		}

		j := 0
		for ; j < len(words[i]); j++ {
			c := words[i][j]
			if m[c - 97] == 0 {
				break;
			}
			m[c - 97]--
		}

		if j == len(words[i]) {
			res += j
		}
	}
	return res
}
