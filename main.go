package main

import "fmt"

func main() {
	s := "A"
	m := [128]int{}

	for i := range s {
		fmt.Println(s[i])
		m[s[i]]++
	}
}


func distributeCandies(candies int, num_people int) []int {

	loop, result := 1, make([]int, num_people)
	for {
		for i := 0; i < num_people; i++ {
			if candies < loop * (i + 1) {
				result = append(result, candies)
				return result
			}
			result[i] += loop * (i + 1)
			candies -= loop * (i + 1)
		}

		if candies == 0 {
			return result
		}
	}
	return result
}